describe('My First Tests', () => {
  beforeEach(() => {
    cy.visit('/login')
  })
  it('Successful Login', () => {
    cy.get('[name="Email"]').type("TestPerson123")
    cy.get('[name="Password"]').type("TestingTesting123")
    cy.get('[type="submit"]').click()
    //Todo: Check if it goes to profile page once that's set up
  })

  it('Unrecognized Username', () => {
    cy.get('[name="Email"]').type("ThisIsNotMyRightAccount")
    cy.get('[name="Password"]').type("TestingTesting123")
    cy.get('[type="submit"]').click()
    cy.contains('Login Failed')
  })

  it('Real Username, wrong password', ()=>{
    cy.get('[name="Email"]').type("TestPerson123")
    cy.get('[name="Password"]').type("TestingTesting124")
    cy.get('[type="submit"]').click()
    cy.contains('Login Failed')
  })

  it('Wrong passwords several times in a row', ()=>{
    cy.get('[name="Email"]').type("TestPerson123")
    cy.get('[name="Password"]').type("TestingTesting124")
    cy.get('[name="Password"]').type("{backspace}5")
    cy.get('[name="Password"]').type("{backspace}6")
    cy.get('[name="Password"]').type("{backspace}7")
    cy.get('[type="submit"]').click()
    cy.contains('Login Failed') //Planning on having this test for if we make an account lockout plicy after # of attempts
  })
  
})
