describe('Profile Page Tests', () => {
  beforeEach(() => {
    cy.visit('/login')
    cy.get('[id="Email"]').type("User1@gmail.com", {force: true})
    cy.get('[id="Password"]').type("UserOneOne")
    cy.get('[type="submit"]').click()
    cy.wait(1000)
  })
  it('Check Right Page', ()=> {
    cy.url().should('include','profile')
  })
  it('Check all goals', ()=>{
    cy.get('[id="allGoals"]').click()
    cy.url().should('include','goals')
  })
  it('go to settings', ()=>{
    cy.get('[id="setButton"]').click()
    cy.url().should('include','settings')
  })
  it('Other\'s Goals', ()=>{
    cy.visit('user/2/profile')
    cy.get('[id="moreButton"]').click()
    cy.get('[id="moreButton"]').click()
    cy.get('[id="moreButton"]').should('not.exist');
  })
  it('Friend Request', ()=>{
    cy.visit('user/12/profile')
    cy.get('[id="friendButton"]').click()
    cy.contains('Request Pending')
  })
})