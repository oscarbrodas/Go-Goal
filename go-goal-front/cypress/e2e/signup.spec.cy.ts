describe('My First Tests', () => {
  beforeEach(() => {
    cy.visit('/sign-up')
  })
  it('Create an account successfully', () => {
    cy.get('[name="FirstName"]').type('TestOne')
    cy.get('[name="LastName"]').type('Person')
    cy.get('[name="Email"]').type('Test@gmail.com')
    cy.get('[name="Username"]').type('TheTestPerson123')
    cy.get('[name="Password"]').type('TestingTesting123')
    cy.get('[formControlName="signUpButton"]').click()
    cy.contains("Account Created")
  })

  it('Submitting with no info', ()=>{
    cy.get('[formControlName = "signUpButton"]').click()
    cy.contains('Not a valid email address')
  })
  it('Not a valid email', () => {
    cy.get('[name="FirstName"]').type('TestTwo')
    cy.get('[name="LastName"]').type('Person')
    cy.get('[name="Email"]').type('Test')
    cy.get('[name="Username"]').type('TheTestPerson234')
    cy.get('[name="Password"]').type('TestingTesting123')
    cy.get('[formControlName="signUpButton"]').click()
    cy.contains("Not a valid email address")
  })

  it('Username already taken', ()=>{
    cy.get('[name="FirstName"]').type('TestThree')
    cy.get('[name="LastName"]').type('Person')
    cy.get('[name="Email"]').type('Test2@gmail.com')
    cy.get('[name="Username"]').type('TheTestPerson123')
    cy.get('[name="Password"]').type('TestingTesting123')
    cy.get('[formControlName="signUpButton"]').click()
    cy.contains("Account Created")
  })
  it('Insecure password', () => {
    cy.get('[name="FirstName"]').type('TestFour')
    cy.get('[name="LastName"]').type('Person')
    cy.get('[name="Email"]').type('Test3@gmail.com')
    cy.get('[name="Username"]').type('TheTestPerson345')
    cy.get('[name="Password"]').type('Testing')
    cy.get('[formControlName="signUpButton"]').click()
    cy.contains("needs a more secure password")
  })

  it('Not a valid email and insecure password', ()=>{
    cy.get('[name="FirstName"]').type('TestFive')
    cy.get('[name="LastName"]').type('Person')
    cy.get('[name="Email"]').type('Test4')
    cy.get('[name="Username"]').type('TheTestPerson456')
    cy.get('[name="Password"]').type('Testing')
    cy.get('[formControlName="signUpButton"]').click()
    cy.contains("Not a valid email address")
  })

  it('Email already taken', ()=>{
    cy.get('[name="FirstName"]').type('TestSix')
    cy.get('[name="LastName"]').type('Person')
    cy.get('[name="Email"]').type('Test@gmail.com')
    cy.get('[name="Username"]').type('TheTestPerson567')
    cy.get('[name="Password"]').type('Testing123')
    cy.get('[formControlName="signUpButton"]').click()
    cy.contains("Account Created")
  })
})
