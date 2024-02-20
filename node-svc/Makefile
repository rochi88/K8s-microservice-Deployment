IMAGE_NAME:= node-svc
TAG:= latest

docker-login:
	@ docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD}

docker-build:
	@ docker build -t ${IMAGE_NAME}:${TAG} .

docker-push: docker-login
	@ docker push ${IMAGE_NAME}:${TAG}