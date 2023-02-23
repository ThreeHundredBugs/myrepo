node {
    checkout scm

    stage('Build image') {
        def serverImage = docker.build("myserver:${env.BUILD_ID}", "./src")
        // serverImage.push()
    }

    stage('Run tests') {
        serverImage.inside {
            sh 'echo Run tests'
            sh 'echo ...'
        }
    }

    stage('Deploy') {
        // sh 'IMAGE=myserver:${env.BUILD_ID} kubectl apply -k kustomize'
    }
}
