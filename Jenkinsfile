node {
    checkout scm

    def serverImage = "myserver:${env.BUILD_ID}"

    stage('Build image') {
        sh "docker build --tag ${serverImage} ./src"
        // docker push
    }

    stage('Run tests') {
        sh "docker run --rm ${serverImage} echo 'run tests....' "
    }

    if (env.BRANCH_NAME == "master") {
        stage('Deploy') {
            // sh 'IMAGE=myserver:${env.BUILD_ID} kubectl apply -k kustomize'
        }
    }
}
