describe('Navigation from User Home Page Tests', () => {
    beforeEach(() => {
      cy.visit('/login')
      cy.get('[name="Email"]').type("sglickman611@gmail.com")
      cy.get('[name="Password"]').type("ThisIsAPassword")
      cy.get('[type="submit"]').click()
      cy.visit('/user/1/profile')
    })
  
    it('Visit profile page from top bar', ()=>{
      cy.get('[id="menuButton"]').click()
      cy.get('[name="profile"]').click()
      cy.url().should('include','profile')
    })
    it('Visit goal page from top bar', ()=>{
      cy.get('[id="menuButton"]').click()
      cy.get('[name="myGoals"]').click()
      cy.url().should('include','goals')
    })
    it('Visit settings page from top bar', ()=>{
      cy.get('[id="menuButton"]').click()
      cy.get('[name="settings"]').click()
      cy.url().should('include','settings')
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