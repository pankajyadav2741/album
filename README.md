# album
Creating a simple REST API to create/delete albums and create/delete images in the album


ENDPOINTS:
==========

#Show all albums
GET /

#Create a new album
POST /{album}

#Delete an existing album
DELETE /{album}

#Show all images in an album
GET /{album}

#Show a particular image inside an album
GET /{album}/{image}

#Create an image in an album
POST /{album}/{image}

#Delete an image in an album
DELETE /{album}/{image}
