describe('Navigation from User Home Page Tests', () => {
     beforeEach(() => {
      cy.visit('/login')
      cy.get('[id="Email"]').type("User1@gmail.com", {force: true})
      cy.get('[id="Password"]').type("UserOneOne")
      cy.get('[type="submit"]').click()
      cy.visit("/user/1/settings")
    })
    it('Visits the initial home page', () => {
      cy.url().should('include','settings')
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
    it('Visit discover page from top bar', ()=>{
      cy.get('[id="menuButton"]').click()
      cy.get('[name="discover"]').click()
      cy.url().should('include','discover')
    })
    it('Visit settings page from top bar', ()=>{
      cy.get('[id="menuButton"]').click()
      cy.get('[name="settings"]').click()
      cy.url().should('include','settings')
    })
    it('Visit Home from bottom bar', () => {
      cy.get('[name = "linkBarHome"]').click()
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
  
    it('Visit profile page from bottom bar', ()=>{
      cy.get('[name="linkBarProfile"]').click()
      cy.url().should('include','1/profile')
    })
  
    it('Logout from bottom bar', ()=>{
      cy.get('[name = "linkBarLogout"]').click()
      cy.url().should('include','main')
    })
})