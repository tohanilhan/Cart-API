
# Patika.dev PropertyFinder Go Bootcamp Final Project

*This repository contains the necessary codes &amp; info about the project given in the Patika.dev PropertyFinder Go bootcamp.*

---

# About the project

*The functions of this service are as follows;*

1. List Products
   - Users should be able to list all products.
2. Add To Cart
   - Users can add their products to the basket and the total of the basket changes accordingly.
3. Show Cart
   - Users can list the products they have added to their cart and total price and VAT of the cart.
4. Delete Cart Item
   - Users can remove the products added from the cart. Notice removing an item may change discount.
5. Complete Order
   - Users can create an order with the products they add to their cart. There is no need to implement any payment mechanism. You can assume that when an order is completed it is fully paid.

*Some business rules*

1. Products always have price and VAT (Value Added Tax, or KDV). VAT might be different for different products. Typical VAT percentage is %1, %8 and %18. So use these values for your products.
2. There might be discount in following situations:
   1.  Every fourth order whose total is more than given amount may have discount
depending on products. Products whose VAT is %1 don’t have any discount but products whose VAT is %8 and %18 have discount of %10 and %15 respectively.
   2. If there are more than 3 items of the same product, then fourth and subsequent ones would have %8 off.
   3. c. If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off.
   4. Only one discount can be applied at a time. Only the highest discount should be applied.

---

### Prerequisites and Installation

- #### **Docker:** 
  You can install Docker Desktop by following the instructions on the [Docker Desktop website](https://desktop.docker.com/).

  If you are using **MacOS**, you can install Docker by following commands:
    - *`brew install docker`*
    - *`docker run hello-world`*
  
 
- #### **PostgreSQL:**

  Plese refer [**here**](https://github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/blob/main/db/README.md) for the initialization scripts.

- #### **Go:**
  
  If you haven't done already, You need to install Go by following the instructions on the [Go website](https://golang.org/doc/install).

- #### **Project:**
  
  You have to clone the project from the [Github repository](https://github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project) and run the below script to build the project
   
   *`git clone https://github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project.git`*

- #### **Run the Project:**

     You have to make sure that you are in the **same directory as the project**. Then you can run the below script:
    *`docker build -t {your-desired-image-name} .`*

     *`docker run -d --env-file=.env --name {your-desired- container-name} -p {your-desired-port}:8080 {your-desired-image-name}`*
   
  Or you can run the project directly by running the below script from the same directory as the project:
    `go run main.go`

- #### **For Testing the Project's endpoints via Postman**
 
  There is a [Postman collection](https://github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/blob/main/PropertyFinder-FinalProject.postman_collection.json) that you can import to Postman and run the tests.

---