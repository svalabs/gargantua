# Setup

1. Secret <code> kubectl apply -f mongodb-secrets.yaml </code>
2. PersistentVolumeClaim <code> kubectl create -f mongodb-pvc.yaml </code>
3. Deployment <code> kubectl apply -f mongodb-deploy.yaml </code>
4. Nodeport Service <code> kubectl create -f mongodb-nodeport-svc.yaml </code>
5. Client <code> kubectl create -f mongodb-client.yaml </code>

# Connection to Client
1. <code>kubectl exec deployment/mongo-client -n hobbyfarm -it -- /bin/bash</code>
2. <code> mongosh --host mongo-service --port 27017 -u adminuser -p password123 </code>

# DB Service
go get -u github.com/gorilla/mux