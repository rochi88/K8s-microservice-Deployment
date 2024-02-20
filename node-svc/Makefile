IMAGE_NAME:= node-svc
TAG:= latest

docker-login:
	@ docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD}

docker-build:
	@ echo ${IMG}
	@ docker build -t ${IMAGE_NAME}:${TAG} .

docker-push:
	@ docker push ${IMAGE_NAME}:${TAG}