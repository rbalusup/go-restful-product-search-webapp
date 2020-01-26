import React from 'react';
import ResultsTable from './ResultsTable';

const API = 'http://localhost:8080/api/v1/products/scan?q=';
let DEFAULT_QUERY = 'Red';
let search = window.location.search;
let params = new URLSearchParams(search);
let searchText = params.get('query');

export default class ConnectedResultsTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            products: []
        };
    }

    componentDidMount() {
        DEFAULT_QUERY = searchText != null && searchText.length > 1 ? searchText.toString() : '';
        this.setState({ isLoading: true });
        fetch(API + DEFAULT_QUERY)
            .then((response) => response.json())
            .then(datum => {
                const results = datum.data.map(result => {
                    return {
                        tcin: result.tcin,
                        title: result.title,
                        price: result.price
                    }
                });
                this.setState({ products: results }).bind(this)
            })
            .catch(e => {
                console.log(e);
                this.setState({...this.state, isFetching: true});
            });
    }
    render() {
        return <ResultsTable results={this.state.products} />;
    }
}
