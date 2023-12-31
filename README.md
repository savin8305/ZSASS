﻿# Car Garage Management API

This project is a simple HTTP (REST) API built using GoFr, meeting the specified requirements.

## Features

- **CRUD Operations:** APIs for create, read, update, and delete operations for car entities.
- **DB Integration:** Integration with a database for data persistence. Supports any SQL or NoSQL DB with a freely available Docker image.
- **Unit Tests:** Achieves a minimum of 60% unit test coverage.
- **Real-world Scenario:** Example scenario of a Car Garage Management service with the following functionalities:
  - Add entry to DB when a car enters the garage.
  - See the list of cars currently in the garage.
  - Update the entry in the DB when car repair is being done, completed, etc.
  - Delete the entry from the DB when the car leaves the garage.

## Bonus Features

The project/repository includes the following bonus items:

- **Unit Test Coverage:** Aiming for > 90% unit test coverage.
- **Postman Collection:** Included for trying out the APIs.

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/savin8305/ZSASS.git
   cd ZSASS
## Sequence Diagram

![Sequence Diagram](https://github.com/savin8305/ZSASS/assets/118232727/02867242-2c89-4d67-aa51-e32c4b84e370)



## POSTMAN APIS

```
{
	"info": {
		"_postman_id": "ddd59b90-7704-4c0a-9bd9-ebf51b625243",
		"name": "Car Garage Management API",
		"schema": "https://schema.getpostman.com/json/collection/v2.0.0/collection.json",
		"_exporter_id": "23774632"
	},
	"item": [
		{
			"name": "Create Car Entry",
			"request": {
				"method": "POST",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n  \"brand\": \"Toyota\",\n  \"model\": \"Camry\",\n  \"year\": 2022,\n  \"condition\": \"New\"\n}"
				},
				"url": {
					"raw": "http://localhost:8080/cars",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"cars"
					]
				}
			},
			"response": [
				{
					"name": "Create Car Entry - Success",
					"originalRequest": {
						"url": {
							"raw": "http://localhost:8080/cars",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"cars"
							]
						},
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"brand\": \"Toyota\",\n  \"model\": \"Camry\",\n  \"year\": 2022,\n  \"condition\": \"New\"\n}"
						}
					},
					"status": "201 Created",
					"code": 201,
					"_postman_previewlanguage": "json",
					"header": [
						{
							"key": "Content-Type",
							"value": "application/json"
						}
					],
					"body": "{\n  \"id\": 1,\n  \"brand\": \"Toyota\",\n  \"model\": \"Camry\",\n  \"year\": 2022,\n  \"condition\": \"New\"\n}",
					"type": "text"
				}
			]
		},
		{
			"name": "Get List of Cars",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/cars",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"cars"
					]
				}
			},
			"response": [
				{
					"name": "Get List of Cars - Success",
					"originalRequest": {
						"url": {
							"raw": "http://localhost:8080/cars",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"cars"
							]
						},
						"method": "GET",
						"header": []
					},
					"status": "200 OK",
					"code": 200,
					"_postman_previewlanguage": "json",
					"header": [],
					"body": "[\n  {\n    \"id\": 1,\n    \"brand\": \"Toyota\",\n    \"model\": \"Camry\",\n    \"year\": 2022,\n    \"condition\": \"New\"\n  }\n]",
					"type": "text"
				}
			]
		},
		{
			"name": "Update Car Entry",
			"request": {
				"method": "PUT",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\n  \"brand\": \"Toyota\",\n  \"model\": \"Camry\",\n  \"year\": 2022,\n  \"condition\": \"Used\"\n}"
				},
				"url": {
					"raw": "http://localhost:8080/cars/{92345}",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"cars",
						"{92345}"
					]
				}
			},
			"response": [
				{
					"name": "Update Car Entry - Success",
					"originalRequest": {
						"url": {
							"raw": "http://localhost:8080/cars/{92345}",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"cars",
								"{92345}"
							]
						},
						"method": "PUT",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"brand\": \"Toyota\",\n  \"model\": \"Camry\",\n  \"year\": 2022,\n  \"condition\": \"Used\"\n}"
						}
					},
					"status": "200 OK",
					"code": 200,
					"_postman_previewlanguage": "json",
					"header": [
						{
							"key": "Content-Type",
							"value": "application/json"
						}
					],
					"body": "{\n  \"id\": {92345},\n  \"brand\": \"Toyota\",\n  \"model\": \"Camry\",\n  \"year\": 2022,\n  \"condition\": \"Used\"\n}",
					"type": "text"
				}
			]
		},
		{
			"name": "Delete Car Entry",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/cars/{92345}",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"cars",
						"{92345}"
					]
				}
			},
			"response": [
				{
					"name": "Delete Car Entry - Success",
					"originalRequest": {
						"url": {
							"raw": "http://localhost:8080/cars/{92345}",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"cars",
								"{92345}"
							]
						},
						"method": "DELETE",
						"header": []
					},
					"status": "204 No Content",
					"code": 204,
					"_postman_previewlanguage": "json",
					"header": [],
					"body": "",
					"type": "text"
				}
			]
		}
	]
}

```
