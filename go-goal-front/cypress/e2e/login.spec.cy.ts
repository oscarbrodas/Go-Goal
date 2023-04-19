describe('My First Tests', () => {
  beforeEach(() => {
    cy.visit('/login')
  })
  it('Successful Login', () => {
    cy.get('[name="Email"]').type("User1@gmail.com")
    cy.get('[name="Password"]').type("UserOne")
    cy.get('[type="submit"]').click()
    //Todo: Check if it goes to profile page once that's set up
    cy.url().should('include','profile')
  })

  it('Unrecognized Username', () => {
    cy.get('[name="Email"]').type("WrongUser@gmail.com")
    cy.get('[name="Password"]').type("UserOneOne")
    cy.get('[type="submit"]').click()
    cy.contains('Login Failed')
  })

  it('Real Username, wrong password', ()=>{
    cy.get('[name="Email"]').type("User1@gmail.com")
    cy.get('[name="Password"]').type("UserOne")
    cy.get('[type="submit"]').click()
    cy.contains('Login Failed')
  })

  it('Wrong passwords several times in a row', ()=>{
    cy.get('[name="Email"]').type("User1@gmail.com")
    cy.get('[name="Password"]').type("TestingTesting124")
    cy.get('[name="Password"]').type("{backspace}5")
    cy.get('[type="submit"]').click()
    cy.get('[name="Password"]').type("{backspace}6")
    cy.get('[type="submit"]').click()
    cy.get('[name="Password"]').type("{backspace}7")
    cy.get('[type="submit"]').click()
    cy.contains('Login Failed') //Planning on having this test for if we make an account lockout policy after # of attempts
  })
  
})
