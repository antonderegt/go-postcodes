# GO Postcodes

## Install
```
git clone https://github.com/antonderegt/go-postcodes.git
```

## Build for Developement (MacBook)
The following command will build the container and hot reload on changes.
```
cd app
docker run -it --rm -w "/go/src/app" -v $(pwd):/go/src/app -p 3000:3000 cosmtrek/air
```

## Build and run for ARM device
```
sudo docker build -t go-pc .
sudo docker run -d --rm -p 3000:3000 go-pc air
```
