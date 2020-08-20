# GO Postcodes

## How To Use
Before you can use any of the interesting endpoints you need to login. The jwt token you receive should be send with the next request. 

### Login
**Endpoint:** (POST) /login

**Parameters:** user, pass

**Example:** curl -d "user=john&pass=doe" http://localhost:3000/login

The returned access token is valid for 72 hours.

### Get Latitude and Longitude
**Endpoint:** (GET) /api/latlon

**Parameters:** street, num, city, postcode (use any combination that makes sense)

**Headers:** Authorization: Bearer \<ACCESS_TOKEN\>

**Example:** curl -H "Authorization: Bearer \<ACCESS_TOKEN\>" http://localhost:3000/api/latlon?street=damrak&num=1&city=amsterdam

### Get Address Details
**Endpoint:** (GET) /api/address

**Parameters:** street, num, city, postalcode (use any combination that makes sense)

**Headers:** Authorization: Bearer \<ACCESS_TOKEN\>

**Example:** curl -H "Authorization: Bearer \<ACCESS_TOKEN\>" http://localhost:3000/api/postcode?street=damrak&num=1&city=amsterdam

### Check Entered Address
**Endpoint:** (GET) /api/check/address

**Parameters:** street, num, city

**Example:** curl http://localhost:3000/api/address?street=damrak&num=1&city=amsterdam

## Prerequisites
A running [Nominatim](https://nominatim.openstreetmap.org/ui/search.html) database

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
