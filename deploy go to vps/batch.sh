git pull
sudo docker build -t golang_health_image .
sudo docker stop golang_health_container
sudo docker rm golang_health_container
sudo docker run -d --name golang_health_container -p 9092:9092 golang_health_image