POST /users/register <br>
`{
  "email" : "non@mail.com",
  "fullname" : "Full",
  "password" : "Testing123",
  "role" : "admin"
}`

POST /users/login <br>
`{
  "email" : "non@mail.com",
  "password" : "Testing123"
}`

POST /products/create  <br>
`{
  "title" : "test",
  "description" : "desc"
}`

GET /product/:product_id

GET /products

UPDATE /products/:product_id <br>
`{
  "title" : "teste",
  "description" : "desc"
}`

DELETE /products/:product_id
