## Backend API
### User
- User: the structure of the User object. The variable names are case sensitive

  - ID: unique id of the user
  
  - CreatedAt: *do not worry about this* Time and date the user was created at
  
  - UpdatedAt: *do not worry about this* Time and date the user was updated at
  
  - DeletedAt: *do not worry about this* Time and date the user was deleted at
  
  - Username: username
  
  - FirstName: first name
  
  - LastName: last name
  
  - Email: the email is unique to each user
  
  - Password: password

  - Description: user set description of their profile

  - XP: experience points (int)
  
- __CreateUser()__ - __POST__

  - Route: __/api/users__
  
  - Input: the body must contain user object
  
  - Output: "Successful", "ErrorExist", and "EmailExist" each as a bool detailing the result of the function

- __GetUser()__ - __GET__

  - Route: __/api/users?id=__
  
  - Input: no body is required. The URL must contain the id of the user you are trying to get, such as getting user with id 9, you need to use __"/api/users?id=9"__
  
  - Output: "ThisUser" as a user object, and "ErrorExist" as a bool

- __UpdateUsername__ - __PUT__

  - Route: __/api/users/{id}/username__
  
  - Input: The User ID (In the Route) to be updated and the new "Username" of type string
  
  - Output: "User" back with the updated change 

- __UpdateFirstname__ - __PUT__

  - Route: __/api/users/{id}/firstname__
  
  - Input: The User ID (In the Route) to be updated and "Firstname" of type string
  
  - Output: "User" back with the updated change 

- __UpdateLastname__ - __PUT__

  - Route: __/api/users/{id}/lastname__
  
  - Input: The User ID (In the Route) to be updated and "Lastname" of type string
  
  - Output: "User" back with the updated change 
  
- __UpdateEmail__ - __PUT__

  - Route: __/api/users/{id}/email__
  
  - Input: The User ID (In the Route) to be updated and "Email" of type string
  
  - Output: "User" back with the updated change 
  
 
- __UpdatePassword__ - __PUT__

  - Route: __/api/users/{id}password__
  
  - Input: The User ID (In the Route) to be updated and "Password" of type string
  
  - Output: "User" back with the updated change 
  
  - __AddXP__ - __PUT__

  - Route: __/api/users/{id}/xp__
  
  - Input: The User ID (In the Route) to be updated and "xp" of type int(32bit)
  
  - Output: "User" back with the updated change 
  
  
- __CheckLogin()__ - __GET__

  - Route: __/api/login/{email}/{password}__ 

  - Notes: Email and Password must be correct or the user object returned will be giberish

  - Input: "Email" and "Password" of type string passed in the route.

  - Output: "FindEmail" and "FindPassword" of type bool. "ThisUser" which is the user object

- __CheckUsername()__ - __GET__

  - Route: __/api/login/{username}__ 

  - Input: "Username" of type string passed in the route.

  - Output: "Exists" and "ValidName", boolean values to validate the "Username" with.


### Friend
- Friend: the structure of the Friend object. The variable names are case sensitive

  - ID: *do not worry about this* Unique id of the relationship
  
  - CreatedAt: *do not worry about this* Time and date the user was created at
  
  - UpdatedAt: *do not worry about this* Time and date the user was updated at
  
  - DeletedAt: *do not worry about this* Time and date the user was deleted at
  
  - User1: id of User1
  
  - User2: id of User2
  
  - WhoSent: int value. 0 means User1 and User2 are friends, 1 means User1 sent a request to User2, 2 means User2 sent a request to User1
  
- __GetAllFriends()__ - __GET__

  - Route: __/api/friends{id}__
  
  - Input: User ID is passed in the route for the user retrieving friends from.
  
  - Output: "IDs" as an array of ints which represent the IDs of the friends of this user. "ErrorExist" as bool
  
- __SendFriendRequest()__ - __PUT__

  - Notes: cannot sent a friend request if you already sent one, are already friends, or the other person sent you a request
  
  - Route: __/api/friends/sendFriendRequest/{sender}/{reciever}__
  
  - Input: Sender's ID and Reciever's ID passed in the route
  
  - Output: "Successful", "ErrorExist" as bools
  
- __GetOutgoingFriendRequests()__ - __GET__

  - Notes: gets the list of people the user sent a request to
  
  - Route: __/api/friends/getOutgoingFriendRequests/{id}__
  
  - Input: User ID passed in route.
  
  - Output: "IDs" as an array of user ids that you sent a request to. "ErrorExist" as a bool
  
- __GetIngoingFriendRequests()__ - __GET__

  - Notes: gets a list of people that sent the user a request
  
  - Route: __/api/friends/getIngoingFriendRequests/{id}__
  
  - Input: User ID passed in route.
  
  - Output: "IDs" as an array of user ids that sent a request to you. "ErrorExist" as a bool
  
- __AcceptFriendRequest()__ - __PUT__

  - Route: __/api/friends/acceptFriendRequest/{sender}/{reciever}__
  
  - Input: Sender/Reciever Pair of IDs passed in route.

  - Output: "Successful", "ErrorExist" as bools.

- __DeclineFriendRequest()__ - __DELETE__

  - Route: __/api/friends/declineFriendRequest/{sender}/{decliner}__
  
  - Input: Sender/Decliner pair of IDs passed in route.

  - Output: "Successful", "ErrorExist" as bools

- __RemoveFriend()__ - __DELETE__

  - Route: __/api/friends/removeFriend/{remover}/{friend}__
  
  - Input: Remover/Friend pair of IDs passed in route.

  - Output: "Successful", "ErrorExist" as bools

### Goal
- Goal: the structure of the Goal object. The variable names are case sensitive

  - ID: *do not worry about this* Unique id of the relationship
  
  - CreatedAt: *do not worry about this* Time and date the user was created at
  
  - UpdatedAt: *do not worry about this* Time and date the user was updated at
  
  - DeletedAt: *do not worry about this* Time and date the user was deleted at
  
  - Title: title of the goal (String)
  
  - Description: description of the goal (String)
  
  - UserID: id of user who has the goal
 
- Benchmark: the structure of the Benchmark object. The variable names are case sensitive

  - ID: *do not worry about this* Unique id of the relationship
  
  - CreatedAt: *do not worry about this* Time and date the user was created at
  
  - UpdatedAt: *do not worry about this* Time and date the user was updated at
  
  - DeletedAt: *do not worry about this* Time and date the user was deleted at
  
  - Description: description of the benchmark (String)
  
  - Completed: Is the benchmark completed? (Bool)

  - GoalID: ID of the goal the benchmark belongs to.

- __CreateGoal()__ - __POST__

  - Route: __/api/goal/{id}__
  
  - Input: Route holds user id and the body take the goal structure
  
  - Output: "Successful", "ErrorExist", and the created Goal. 

- __GetGoals()__ - __GET__

  - Route: __/api/goals/{id}__
  
  - Input: User ID is passed through the route.
  
  - Output: "Successful", "ErrorExist", and a list of goals is returned.

- __DeleteGoal()__ - __DELETE__

  - Route: __/api/goals/{id}__
  
  - Input: Goal ID is passed through the route.
  
  - Output: "Successful", "ErrorExist" are returned.

- __AddBenchmark()__ - __POST__

  - Route: __/goals/{id}/{goalID}__
  
  - Input: Goal ID is passed through the route. The body of request contains the input benchmark object with only the description required.
  
  - Output: "Successful", "ErrorExist" are returned.

- __GetBenchmarks()__ - __GET__

  - Route: __/goals/benchmarks/{goalID}__
  
  - Input: Goal ID is passed through the route.
  
  - Output: "Successful", "ErrorExist" are returned. "Benchmarks" is an array of benchmark objects

- __UpdateBenchmarkDescription()__ - __PUT__

  - Route: __/goals/benchmarks/description/{benchmarkID}__
  
  - Input: Goal ID is passed through the route.
  
  - Output: "Successful", "ErrorExist" are returned.

- __UpdateBenchmarkCompletion()__ - __PUT__

  - Route: __/goals/benchmarks/completion/{benchmarkID}__
  
  - Input: Goal ID is passed through the route.
  
  - Output: "Successful", "ErrorExist" are returned.

- __DeleteBenchmark()__ - __DELETE__

  - Route: __/goals/benchmarks/{benchmarkID}__
  
  - Input: Goal ID is passed through the route.
  
  - Output: "Successful", "ErrorExist" are returned.

## Backend Sprint 4 Work

Added search friends function in order to search for new friends with a username similar to the one the user will be searching. Updated the unit tests for friends to reflect the proper functionality of friends. Created unit tests for the benchmark functionality. Added profile descriptions to the user table. Updated the routes and functions of the update user functions to accect json objects within the body of the requests. Added AWS s3 functionality, routes, and functions to support profile pictures.

## Frontend Sprint 4 Work

- Created main page for new users to the website to view and explain the service.

- Implemented the Discover page to search, add, and remove friends and users.

- Added the 

- Implemented reactive CSS for all pages.

## Frontend Unit Tests

### Navigation Suite
 - For each individual page in the website, the following tests were done
 - Visit Home from Bottom Bar - Tests if home button properly routes to main page
 - Visit Help from Bottom Bar - Tests if help button properly routes to its page
 - Visit About Us from Bottom Bar - Tests if about button properly routes to its page
 - Visit Login from Bottom Bar - Tests if bottom login button properly routes to login page
 - Visit Sign-Up from Bottom Bar - Tests if bottom sign-up button properly routes to sign-up page

 - Tests of the top bar depended on if a page was part of the user component. Pages that were had these tests:
 - Visit Profile Page from Top Bar - Tests if menu icon properly routes to profile page
 - Visit Goal Page from Top Bar - Tests if menu icon properly routes to goal page
 - Visit Settings Page from Top Bar - Tests if menu icon properly routes to settings page
 - Visit Discover Page from Top Bar - Tests if menu icon properly routes to discover page (currently 404)

 - Other pages had these tests:
 - Visit Login Page from Top Bar - Tests if login button properly routes to login page
 - Visit Sign-Up Page from Top Bar - Tests if sign-up button properly routes to sign-up page
### Sign Up Tests
 - Create an Account Successfully - Put in proper input for every paramter and see if account is correctly created
 - Submitting with no info - Simply hit submit without any user info and ensure it does not allow sign-up
 - Not a valid email - Attempt a sign-up with an email that doesn't have an '@'
 - Username already taken - Check if inputting a username that has already been assigned to another account will allow a sign-up to go through
 - Insecure Password - Check if trying to create a password less than 8 characters will cause an error
 - Insecure Password and not a valid email - Try inputting a submission with both conditions that should pose an error
 - Email Already taken - Try signing up with an email already in the system, which should not work it
### Login Tests
 - Successful Login - Input a correct email and password and see if page redirects
 - Unrecognized Username - Input a username that is not in the system and check if log in fails
 - Real Username, Wrong Password - Use the wrong password and check if log in fails
 - Wrong password several times in a row - Use the wrong password for multiple attempts to test for a lockout policy
### Profile Tests
 - Check Right Page - Ensure navigation leads to correct page
 - Check All Goals - Use "go to all goals" link to go to goals page
 - Go to settings - Use "go to account settings" link to go to settings page
 - Other's Goals - Go to another's profile and test if more button brings up more goals
 - Friend Request - Go to another's profile and test if friend request button will give request pending message
### Settings Tests
 - Change First Name - Test if input field will allow you to edit first name
 - Change Last Name - Test if input field will allow you to edit last name
 - Change Email - Test if input field will allow you to edit email
 - Change Username - Test if input field will allow you to edit username
 - No Changes - Test if cancel button will not allow changes to go through
 - Invalid Email - Test if editing email to invalid format will not allow the edit to go through
 - Bad Password - Test if editing password to insecure password will not allow the edit to go through
### Goals Tests
 - Add Goal - Type in values and add goal to a system
 - Persistent Goal - check if a new goal stays on the page after refreshing and moving to and from that tab
 - No Goal - check if an exception is raised when a goal is entered with no data
 - Complete Goal - Click complete and see if that info is saved when reloaded
 - Delete Goal - Click delete options and see if goal is gone when reloaded
 - Add Benchmark - Tests to see if a benchmark was properly added to the database with the proper foreign key to the appropriate goals object
 - Get Benchmarks - tests proper getting of benchmarks, checks if no benchmark is returned when no benchmark exists, and tests for when input goal ID does not exist
