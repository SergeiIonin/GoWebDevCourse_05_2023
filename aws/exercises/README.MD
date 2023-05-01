
1) change port number to 80
2)  GOOS=linux GOARCH=amd64 go build
    (for linux machine, ubuntu e.g.)
3) scp -i ~/.ssh/go_ec2_kp_1.pem ./exercises ubuntu@ec2-18-212-36-127.compute-1.amazonaws.com:
   (dns will vary)
4similarly, push template files to EC2 instance
5) run the app, enable it via systemd
ssh -i ~/.ssh/go_ec2_kp_1.pem ubuntu@ec2-18-212-36-127.compute-1.amazonaws.com