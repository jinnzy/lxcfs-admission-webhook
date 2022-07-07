#!/bin/bash

service='lxcfs-admission-webhook'
namespace='default'
tmpdir=$(mktemp -d)

echo "creating certs in tmpdir ${tmpdir} "


openssl genrsa -out ${tmpdir}/ca.key 2048

openssl req -sha256 -new -x509 -days 10950 -key ${tmpdir}/ca.key -subj "/CN=${service}" -out ${tmpdir}/ca.crt

openssl req -sha256 -newkey rsa:2048 -nodes -keyout ${tmpdir}/tls.key -subj "/CN=${service}.${namespace}.svc" -out ${tmpdir}/tls.csr

echo "subjectAltName=DNS:${service}.${namespace}.svc" > ${tmpdir}/extensions.txt
echo "extendedKeyUsage=serverAuth" >> ${tmpdir}/extensions.txt

openssl x509 -sha256 -req -extfile ${tmpdir}/extensions.txt -days 10585 -in ${tmpdir}/tls.csr -CA ${tmpdir}/ca.crt -CAkey ${tmpdir}/ca.key -CAcreateserial -out ${tmpdir}/tls.crt

echo
echo ">> Generating kube secrets..."
kubectl create secret tls ${service}-tls \
  --cert=${tmpdir}/tls.crt \
  --key=${tmpdir}/tls.key \
  --dry-run=client -o yaml |
  kubectl -n ${namespace} apply -f -


kubectl get secret lxcfs-admission-webhook-tls

kubectl create -f deployment/deployment.yaml
kubectl create -f deployment/service.yaml

CA_BUNDLE=$(cat ${tmpdir}/ca.crt | base64 -w 0)
sed -e "s#\${CA_BUNDLE}#${CA_BUNDLE}#g" ./deployment/mutatingwebhook.yaml  > ./deployment/mutatingwebhook-ca-bundle.yaml


kubectl create -f deployment/mutatingwebhook-ca-bundle.yaml

