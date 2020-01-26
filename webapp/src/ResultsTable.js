import React from 'react';
import { Table, Header, Segment } from 'semantic-ui-react'

export default function ResultsTable({results}) {

    // this._onChange = this._onChange.bind(this)

    const rows = results.map(((result, index) => {
        return (
            <Table.Row key={ index }>
                <Table.Cell>{ result.tcin }</Table.Cell>
                <Table.Cell>{ result.title }</Table.Cell>
                <Table.Cell>{ result.price }</Table.Cell>
            </Table.Row>
        );
    }));

    return (
            <div className="ui container">
                <Segment>
                    <Header>Products </Header>
                    <Table striped>
                        <Table.Header>
                            <Table.Row>
                                <Table.HeaderCell>TCIN</Table.HeaderCell>
                                <Table.HeaderCell>Title</Table.HeaderCell>
                                <Table.HeaderCell>Price</Table.HeaderCell>
                            </Table.Row>
                        </Table.Header>
                        <Table.Body>
                            { rows }
                        </Table.Body>
                    </Table>
                </Segment>
            </div>
    );
}
