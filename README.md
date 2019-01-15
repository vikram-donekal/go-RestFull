# go-RestFull

This example is about RestFull Web service in Go Lang:

weservcies:
	/api/findall- GET 
	/api/create - POST 
	/api/update - PUT
	/api//delete/{id} - DELETE 
	/api/exit - GET 
	
	
Example also includes LoggingHandler which will Log Each and every Service call in a proper fashion.


How to Start:
 1:Using Docker Image
	
		docker run -d -p 9090:9090 vikramdonekal/go-restfull 
		
		Test with:  http://machineIp:9090/api/exit
	
 2:Using Helm Chart inside K8
	Download the helm chart from my github		
		helm package /path/to/my-helm-chart
		helm install /output/of/package.tgz
	Verify Installion:
		kubectl get pods 
		kubectl get services
Get Port of the service mapping and verify with : http://machineIp:servicePort:/api/exit 
		
		
		
		
