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
	stage('unit tests') {
	    steps {
		sh 'go install github.com/jstemmer/go-junit-report@latest'
		sh 'cd hypernyms ; go test -v vocab | ~/go/bin/go-junit-report > ../unit-tests.xml'
	    }
	}	
    }
    post {
	always {
	    archiveArtifacts artifacts: 'hypernyms/hypernyms', fingerprint: true
	    junit 'unit-tests.xml'
	}
    }
}
