describe('My First Tests', () => {
  beforeEach(() => {
    cy.visit('/')
  })
  it('Visits the initial project page', () => {
    cy.contains('Carousel')
  })

  it('Visit login page and type info', ()=>{
    cy.get('[formControlName = "login"]').click()
  })
  
})
