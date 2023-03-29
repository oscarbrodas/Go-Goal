import { NavbarTopComponent } from '../src/app/navbar/navbar-top/navbar-top.component';
describe('navbar.cy.ts', () => {
  beforeEach(() => {
    cy.mount(NavbarTopComponent)
  })

  it('main page link', () => {
    cy.get('[src="assets\gg.png"]').click()
    cy.url().should('include','main')
  })
  
  it('login page link', () => {
    cy.get('[id = "navBarLogin"]').click()
    cy.url().should('include','login')
  })

  it('sign up page link', () => {
    cy.get('[id = "navBarSignUp"]').click()
    cy.url().should('include','sign-up')
  })
})