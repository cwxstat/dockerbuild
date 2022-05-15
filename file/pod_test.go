package file

import (
	"testing"
)

var data = `
apiVersion: v1
kind: Pod
metadata:
  name: active-incident
  labels:
    app.kubernetes.io/name: active-incident
spec:
  volumes:
    - name: active-incident-data-volume
      persistentVolumeClaim:
        claimName: active-incident-storage
        readOnly: false
  containers:
    - name: active-incident-container
      image: us-central1-docker.pkg.dev/mchirico/public/activeincident:v0.1.3
      env:
      - name: MONGO_URI
        valueFrom:
          secretKeyRef:
            name: mongo-creds
            key: mongouri

      - name: OWM_API_KEY
        valueFrom:
          secretKeyRef:
            name: mongo-creds
            key: weatherkey

      - name: MONGO_WEATHER
        valueFrom:
          secretKeyRef:
            name: mongo-creds
            key: weathercollection


      volumeMounts:
        - mountPath: /etc/mongo
          name: active-incident-data-volume
  restartPolicy: Always

`

func TestFindImages(t *testing.T) {

	result, err := FindImages(data)
	if err != nil {
		t.Fatal(err)
	}

	if result[0] != "us-central1-docker.pkg.dev/mchirico/public/activeincident:v0.1.3" {
		t.Fatalf("bad result")
	}

}
