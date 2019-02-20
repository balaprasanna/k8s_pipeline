# Steps to build

1. Build the image
```
docker build -t balanus/newsapp:nodejs-v1 .
```

2. Run the image
```
docker run -it -p 3000:3001 -e PORT=3001 --rm balanus/newsapp:nodejs-v1
```

3. Push the image
```
docker login
docker push balanus/newsapp:nodejs-v1
```
