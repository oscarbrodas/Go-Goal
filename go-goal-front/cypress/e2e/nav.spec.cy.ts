describe('Navigation from Main Page Tests', () => {
  beforeEach(() => {
    cy.visit('/')
  })
  it('Visits the initial project page', () => {
    cy.contains('Carousel')
  })

  it('Visit login page from top bar', ()=>{
    cy.get('[id = "navBarLogin"]').click()
    cy.url().should('include','login')
  })
  it('Visit sign-up page from top bar', ()=>{
    cy.get('[id = "navBarSignUp"]').click()
    cy.url().should('include','sign-up')
  })
  it('Visit Home from bottom bar', () => {
    cy.get('[name = "linkBarMain"]').click()
    cy.url().should('include','main')
  })

  it('Visit Help from bottom bar', ()=>{
    cy.get('[name = "linkBarHelp"]').click()
    cy.url().should('include','help')
  })

  it('Visit About Us from bottom bar', ()=>{
    cy.get('[name = "linkBarAboutUs"]').click()
    cy.url().should('include','aboutus')
  })

  it('Visit login page from bottom bar', ()=>{
    cy.get('[name = "linkBarLogin"]').click()
    cy.url().should('include','login')
  })

  it('Visit sign-up page from bottom bar', ()=>{
    cy.get('[name = "linkBarSignUp"]').click()
    cy.url().should('include','sign-up')
  })
})

describe('Navigation from Login Page Tests', () => {
  beforeEach(() => {
    cy.visit('/login')
  })
  it('Visits the initial project page', () => {
    cy.url().should('include','login')
  })

  it('Visit login page from top bar', ()=>{
    cy.get('[id = "navBarLogin"]').click()
    cy.url().should('include','login')
  })
  it('Visit sign-up page from top bar', ()=>{
    cy.get('[id = "navBarSignUp"]').click()
    cy.url().should('include','sign-up')
  })
  it('Visit Home from bottom bar', () => {
    cy.get('[name = "linkBarMain"]').click()
    cy.url().should('include','main')
  })

  it('Visit Help from bottom bar', ()=>{
    cy.get('[name = "linkBarHelp"]').click()
    cy.url().should('include','help')
  })

  it('Visit About Us from bottom bar', ()=>{
    cy.get('[name = "linkBarAboutUs"]').click()
    cy.url().should('include','aboutus')
  })

  it('Visit login page from bottom bar', ()=>{
    cy.get('[name = "linkBarLogin"]').click()
    cy.url().should('include','login')
  })

  it('Visit sign-up page from bottom bar', ()=>{
    cy.get('[name = "linkBarSignUp"]').click()
    cy.url().should('include','sign-up')
  })
})

describe('Navigation from Sign Up Page Tests', () => {
  beforeEach(() => {
    cy.visit('/sign-up')
  })
  it('Visits the initial project page', () => {
    cy.url().should('include','sign-up')
  })

  it('Visit login page from top bar', ()=>{
    cy.get('[id = "navBarLogin"]').click()
    cy.url().should('include','login')
  })
  it('Visit sign-up page from top bar', ()=>{
    cy.get('[id = "navBarSignUp"]').click()
    cy.url().should('include','sign-up')
  })
  it('Visit Home from bottom bar', () => {
    cy.get('[name = "linkBarMain"]').click()
    cy.url().should('include','main')
  })

  it('Visit Help from bottom bar', ()=>{
    cy.get('[name = "linkBarHelp"]').click()
    cy.url().should('include','help')
  })

  it('Visit About Us from bottom bar', ()=>{
    cy.get('[name = "linkBarAboutUs"]').click()
    cy.url().should('include','aboutus')
  })

  it('Visit login page from bottom bar', ()=>{
    cy.get('[name = "linkBarLogin"]').click()
    cy.url().should('include','login')
  })

  it('Visit sign-up page from bottom bar', ()=>{
    cy.get('[name = "linkBarSignUp"]').click()
    cy.url().should('include','sign-up')
  })
})
