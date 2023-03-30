describe('Goal Page Tests', () => {
  beforeEach(() => {
    cy.visit('/login')
    cy.get('[name="Email"]').type("sglickman611@gmail.com")
    cy.get('[name="Password"]').type("ThisIsAPassword")
    cy.get('[type="submit"]').click()
    cy.visit('user/1/goals')
  })
  it('Add goal', ()=>{
    cy.get('[name="Title"]').type("goal1")
    cy.get('[name="Description"]').type("description1")
    cy.get('[name="submit"]').click()
    cy.contains("goal1")
    
  })
  it('Persistent goal', ()=>{
    cy.get('[name="Title"]').type("goal2")
    cy.get('[name="Description"]').type("description2")
    cy.get('[name="submit"]').click()
    cy.visit('user/1/profile')
    cy.visit('user/1/goals')
    cy.contains("goal1")
    
  })
  it('No goal', ()=>{
    cy.get('[name="submit"]').click()
    cy.contains("")
  })
  it('Complete Goal', ()=>{
    cy.get('[class="complete-b"]').click()
    cy.reload()
  })
  it("Delete Goal", ()=>{
    cy.get("[name='delete']").click()
    cy.get('[class="complete-b"]').click()
    cy.reload()
    cy.get('[class="complete-b"]').should('not-exist')
  })
})