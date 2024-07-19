export DEPLOY_IMG_TAG=0.0.1

# 定义变量
IMAGE_NAME="mythosma/go-server"
TAG=$DEPLOY_IMG_TAG # 可以是 latest，或者特定的版本号，如 0.0.1

# 步骤 1: 构建 Docker 镜像
echo "Building Docker image..."
docker build -t ${IMAGE_NAME}:${TAG} .

# 步骤 2: 推送镜像到 Docker Hub
echo "Pushing image to Docker Hub..."
docker push ${IMAGE_NAME}:${TAG}