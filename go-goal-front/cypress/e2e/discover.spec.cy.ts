describe('Discover Page Tests', () => {
    beforeEach(() => {
      cy.visit('/login')
      cy.get('[id="Email"]').type("User1@gmail.com", {force: true})
      cy.get('[id="Password"]').type("UserOneOne")
      cy.get('[type="submit"]').click()
      cy.wait(1000)
      cy.visit("/user/1/discover")
      cy.wait(1000)
    })
    it('Check Right Page', ()=> {
      cy.url().should('include','discover')
    })
    it('Check friend', ()=>{
      cy.contains('UserTwo').click({force: true})
      cy.contains('ID: 2')
    })
    it('Search for User', ()=>{
      cy.get('[id="s"]').type("UserEl", {force: true})
      cy.get('[id="searchButton"]').click({force: true})
      cy.contains('Eleven')
    })
    it('Click through Users',()=>{
        cy.get('[id="s"]').type("UserT", {force: true})
      cy.get('[id="searchButton"]').click({force: true})
      cy.get('[id="nextButton"]').click()
      cy.get('[id="nextButton"]').click()
      
      cy.contains('Ten')
    })
    it('Click Backwards',()=>{
        cy.get('[id="s"]').type("UserT", {force: true})
      cy.get('[id="searchButton"]').click({force: true})
      cy.get('[id="nextButton"]').click()
      cy.get('[id="nextButton"]').click()
      cy.get('[id="prevButton"]').click()
      cy.get('[id="prevButton"]').click()
      cy.get('[id="prevButton"]').click()
      cy.get('[id="prevButton"]').click()
      
      cy.contains('Twelve')
    })
    it('Visit Profile',()=>{
        cy.get('[id="s"]').type("UserT", {force: true})
      cy.get('[id="searchButton"]').click({force: true})
      cy.get('[id="nextButton"]').click()
      cy.get('[id="nextButton"]').click()
      cy.get('[id="nextButton"]').click()
      cy.contains('View Profile').click({force: true})
      cy.url().should('include','12/profile')
    })
  })