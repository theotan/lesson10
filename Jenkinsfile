pipeline {
    agent any
    tools {
	go 'Golang version 17.3'
    }
    stages {
	stage('compilation') {
	    steps {
		sh 'cd hypernyms ; go build'
	    }
	}
	stage('smoke-test') {
	    steps {
		sh './hypernyms/hypernyms'
	    }
	}
    }
    post {
	always {
	    archiveArtifacts artifacts: 'hypernyms/hypernyms', fingerprint: true
	}
    }
}
