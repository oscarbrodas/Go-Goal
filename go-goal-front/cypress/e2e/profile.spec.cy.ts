describe('Profile Page Tests', () => {
  beforeEach(() => {
    cy.visit('/login')
    cy.get('[name="Email"]').type("sglickman611@gmail.com")
    cy.get('[name="Password"]').type("ThisIsAPassword")
    cy.get('[type="submit"]').click()
    cy.visit('user/1/profile')
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
    cy.visit('user/8/profile')
    cy.get('[id="moreButton"]').click()
    cy.get('[id="moreButton"]').click()
    cy.get('[id="moreButton"]').should('not.exist');
  })
  it('Friend Request', ()=>{
    cy.visit('user/8/profile')
    cy.get('[id="friendButton"]').click()
    cy.contains('Request Pending')
  })
})