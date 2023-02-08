/// <reference types="cypress" />


describe('Transaction API testing', () =>{
    let transactionItem;
    it('Fetch Transaction items - GET', () => {
        cy.request('/transactions/').as('transactionRequest');
        cy.get('@transactionRequest').then(transactions => {
            expect(transactions.status).to.eq(200);
            assert.isArray(transactions.body, 'Transaction items')    
        })        
    })


    it('deletes Transaction items - DELETE', () => {
        cy.request('DELETE', `/transactions/${transactionItem}`).as('transactionRequest');
        // deletes Todo item with id = 9
        cy.get('@transactionRequest').then(todos => {
            expect(todos.status).to.eq(200);
            assert.isString(todos.body, 'transaction deleted!')
        });
    });
})