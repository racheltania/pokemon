import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import '../PokemonList.css';
import '../MenuBar.css';

const PokemonList = () => {
    const [pokemonList, setPokemonList] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);
    const [limit, setLimit] = useState(10);

    function extractParamFromUrl(url) {
        const match = url.match(/\/(\d+)\/$/);
        const paramValue = match ? match[1] : null;
        return `https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/${paramValue}.png`;
    }

    useEffect(() => {
        fetch(`http://localhost:1323/?page=${currentPage}&limit=${limit}`)
            .then((response) => {
                if (!response.ok) {
                    throw new Error(`Network response was not ok: ${response.statusText}`);
                }
                return response.json();
            })
            .then((data) => {
                setPokemonList(data);
            })
            .catch((error) => console.error('Error fetching data:', error));
    }, [currentPage, limit]);

    const handlePageChange = (newPage) => {
        setCurrentPage(newPage);
    };

    const handleLimitChange = (event) => {
        setLimit(parseInt(event.target.value, 10));
        setCurrentPage(1); // Reset current page when changing the limit
    };

    return (
        <div>
            <div className="menu-bar">
                <Link to="/" className="menu-item">
                    <img src={'/pokeball.png'} alt="Custom Icon" style={{ marginRight: '9px', width: '30px', height: '20px' }} />
                </Link>
                <Link to="/mypokemon" className="menu-item">
                    My Pokemon List
                </Link>
            </div>
            <div className="pokemon-list-container">
                <h1>Explore Pokemon</h1>
                <ul className="pokemon-list">
                    {pokemonList.map((pokemon) => (
                        <li key={pokemon.name} className="pokemon-item">
                            <Link to={`/${pokemon.name}`} className="pokemon-link">
                                <div className="pokemon-card">
                                    <img
                                        src={extractParamFromUrl(pokemon.url)}
                                        alt={pokemon.name}
                                        className="pokemon-image"
                                    />
                                    <h3>{pokemon.name}</h3>
                                </div>
                            </Link>
                        </li>
                    ))}
                </ul>
                <div className="pagination">
                    <button onClick={() => handlePageChange(currentPage - 1)} disabled={currentPage === 1}>
                        Previous
                    </button>
                    <span>Page {currentPage}</span>
                    <button onClick={() => handlePageChange(currentPage + 1)}>Next</button>
                </div>
                <div>
                    <select id="limit" value={limit} onChange={handleLimitChange}>
                        <option value="5">5</option>
                        <option value="10">10</option>
                        <option value="15">15</option>
                        <option value="20">20</option>
                    </select>
                </div>
            </div>
        </div>
    );
};

export default PokemonList;
