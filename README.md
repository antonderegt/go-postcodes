# GO Postcodes

## How To Use

### Get Lat + Lon
**Endpoint:** api/latlon

**Parameters:** street, num, city

**Example:** http://localhost:3000/api/latlon?street=damrak&num=1&city=amsterdam

### Get Postcode
**Endpoint:** api/postcode

**Parameters:** street, num, city

**Example:** http://localhost:3000/api/postcode?street=damrak&num=1&city=amsterdam

### Check Entered Address
**Endpoint:** api/address

**Parameters:** street, num, city

**Example:** http://localhost:3000/api/address?street=damrak&num=1&city=amsterdam

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
cd app/
sudo docker build -t go-pc .
sudo docker run -d --rm -p 3000:3000 go-pc air
```
