# Go-Goal Sprint 2

## Backend Unit Tests

- TestGetAllFriends(): to properly get a list of all of your friends

- TestSendFriendRequest1(): if you can send a friend request to a user

- TestSendFriendRequest2(): to see if error is properly returned after sending an invalid friend request

- TestGetOutgoingFriendRequests(): to get a list of all users that you sent a friend rquest to

- TestGetIngoingFriendRequests(): to get a list of all users that sent you a friend request

- TestAcceptFriendRequest(): to accept a friend request

- TestDeclineFriendRequest(): to decline a friend request

- TestRemoveFriend(): to delete a friend from your friend's list

- TestGetUser(): to get the corresponding user object from the database

- TestCreateUser1(): to create a user where the email is not in use

- TestCreateUser2(): to create a user and return an error when the email is in use

- TestCheckLogin1(): log in when the email and password are correct

- TestCheckLogin2(): log in when the email is correct but password is not

## Backend API
### User
- User: the strructure of the User object. The variable names are case sensitive

  - ID: unique id of the user
  
  - CreatedAt: do not worry about this. Time and date the user was created at
  
  - UpdatedAt: do not worry about this. Time and date the user was updated at
  
  - DeletedAt: do not worry about this. Time and date the user was deleted at
  
  - Username: username
  
  - FirstName: first name
  
  - LastName: last name
  
  - Email: the email is unique to each user
  
  - Password: password
  
- CreateUser(): POST

  - Route: /users
  
  - Input: the body must contain user object
  
  - Output: "Successful", "ErrorExist", and "EmailExist" each as a bool detailing the result of the function

- GetUser(): GET

  - Route: /users?id=
  
  - Input: no body is required. The URL must contain the id of the user you are trying to get, such as getting user with id 9, you need to use "/users?id=9"
  
  - Output: "ThisUser" as a user object, and "ErrorExist" as a bool

- CheckLogin(): GET

  - Notes: Email and Password must be correct or the user object returned will be giberish

  - Input: "Email" and "Password" of type string

  - Output: "FindEmail" and "FindPassword" of type bool. "ThisUser" which is the user object
  
### Friend
- Friend: the strructure of the Friend object. The variable names are case sensitive

  - ID: do not worry about this. Unique id of the relationship
  
  - CreatedAt: do not worry about this. Time and date the user was created at
  
  - UpdatedAt: do not worry about this. Time and date the user was updated at
  
  - DeletedAt: do not worry about this. Time and date the user was deleted at
  
  - User1: id of User1
  
  - User2: id of User2
  
  - WhoSent: int value. 0 means User1 and User2 are friends, 1 means User1 sent a request to User2, 2 means User2 sent a request to User1
  
- GetAllFriends(): GET

  - Route: /friends
  
  - Input: body must contain the user object of the current user
  
  - Output: "IDs" as an array of ints which represent the IDs of the friends of this user. "ErrorExist" as bool
  
- SendFriendRequest(): PUT

  - Notes: cannot sent a friend request if you already sent one, are already friends, or the other person sent you a request
  
  - Route: /friends/sendFriendRequest?id=
  
  - Input: the body must contain the user object of the user sending the request. The URL must contain the id of the user recieving the request
  
  - Output: "Successful", "ErrorExist" as bools
  
- GetOutgoingFriendRequests(): GET

  - Notes: gets the list of people that you sent a request to
  
  - Route: /friends/getOutgoingFriendRequests
  
  - Input: body must contain the user object
  
  - Output: "IDs" as an array of user ids that you sent a request to. "ErrorExist" as a bool
  
- GetIngoingFriendRequests(): GET

  - Notes: gets a list of people that sent you a request
  
  - Route: /friends/getIngoingFriendRequests
  
  - Input: body must contain the user object
  
  - Output: "IDs" as an array of user ids that sent a request to you. "ErrorExist" as a bool
  
- AcceptFriendRequest(): PUT

  - Route: /friends/acceptFriendRequest?id=
  
  - Input: body must contain the user object. The URL must contain the id of the user that sent you the request

  - Output: "Successful", "ErrorExist" as bools

- DeclineFriendRequest(): DELETE

  - Route: /friends/declineFriendRequest?id=
  
  - Input: body must contain the user object. The URL must contain the id of the user that sent you the request

  - Output: "Successful", "ErrorExist" as bools

- RemoveFriend(): DELETE

  - Route: /friends/removeFriend?id=
  
  - Input: body must contain the user object. The URL must contain the id of the user that you are removing

  - Output: "Successful", "ErrorExist" as bools

### FRONTEND TESTS:
 - Sign Up Form Tests
   - Correctly Sign Up: Given valid info for all parameters, ensuring the page adds them to the databse and routes the user to their page
   - Invalid Email: Given an invalid email, ensuring the website does not take those values as an account
   - Insecure Password: Given a password less than 8 characters, not allowing a user to sign up until it is further secured
   - Already Taken Email: Does not allow an account to be signed up if the email is already in the database (THIS TEST EXISTS BUT FUNCTIONALITY IS NOT IMPLEMENTED YET)
   - Already Taken Username: Does not allow an account to be signed up if the username is already in the database (THIS TEST EXISTS BUT FUNCTIONALITY IS NOT IMPLEMENTED YET)
- Login Form Tests:
   - Correctly Login: Login with a currently used email and password and route to the profile page on success
   - Wrong Password: Login with an email that's in the database but a password that is not, get a login failed message
   - Wrong email: Login with an email not in the database and receive a failed login message
   - Multiple Wrong Attempts: Try multiple login attempts on the same email, currently just check it always fails but may potentially add an account lock-out policy and check for that
- Navigation Tests:
   - Main Page Suite: Try every link on the top nav bar and link bar from the main page to other pages (including itself)
   - Login Page Suite: Try every link on the top nav bar and link bar from the login page to other pages (including itself)
   - Sign Up Page Suite: Try every link on the top nav bar and link bar from the sign up page to other pages (including itself)
