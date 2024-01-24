import React, { useEffect, useState } from 'react';
import Modal from "react-modal";
import {Link} from "react-router-dom";
import '../PokemonList.css';
import '../MenuBar.css';

const MyPokemon = () => {
    const [pokemonList, setPokemonList] = useState([]);
    const [showSuccessModal, setShowSuccessModal] = useState(false);
    const [showFailureModal, setShowFailureModal] = useState(false);
    const [showSureModal, setShowSureModal] = useState(false);
    const [selectedPokemonId, setSelectedPokemonId] = useState(null);
    const [selectedPokemonName, setSelectedPokemonName] = useState(null);
    const [selectedPokemonPicture, setSelectedPokemonPicture] = useState(null);
    const [currentPage, setCurrentPage] = useState(1);
    const [limit, setLimit] = useState(10);

    useEffect(() => {
        fetch(`http://localhost:1323/my?page=${currentPage}&limit=${limit}`)
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

    const handleRelease = async () => {
        try {
            setShowSureModal(false)
            if (selectedPokemonId) {
                const response = await fetch(`http://localhost:1323/release/${selectedPokemonId}`, {
                    method: 'DELETE',
                });

                if (!response.ok) {
                    throw new Error(`Failed to release data. Status: ${response.status}`);
                }

                const result = await response.json();
                if (result) {
                    setPokemonList((prevList) => prevList.filter(pokemon => pokemon.id !== selectedPokemonId));
                    setShowSuccessModal(true);
                } else {
                    setShowFailureModal(true);
                }
            }
        } catch (error) {
            console.error('Error releasing Pokemon:', error.message);
        }
    };

    const handleQuestion = (pokemonId,pokemonName,pokemonPicture) => {
        setSelectedPokemonId(pokemonId);
        setSelectedPokemonName(pokemonName)
        setSelectedPokemonPicture(pokemonPicture)
        setShowSureModal(true)
    };

    return (
        <div>
            <div className="menu-bar">
                <Link to="/" className="menu-item">
                    <img src={'/pokeball.png'} alt="Custom Icon"
                         style={{marginRight: '9px', width: '30px', height: '20px'}}/>
                </Link>
            </div>
            <div className="pokemon-list-container">
                <h1>My Pokemon List</h1>
                <ul className="pokemon-list">
                    {pokemonList.map((pokemon) => (
                        <li key={pokemon.id} className="pokemon-item">
                            <div className="pokemon-card">
                                <img
                                    src={pokemon.picture}
                                    alt="Pokemon"
                                    className="pokemon-image"
                                />
                                <h3>{pokemon.name}</h3>
                                <button
                                    onClick={() => handleQuestion(pokemon.id, pokemon.name, pokemon.picture)}>Release
                                </button>
                            </div>
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
                <Modal
                    isOpen={showSureModal}
                    onRequestClose={() => setShowSureModal(false)}
                    contentLabel="Sure Modal"
                    appElement={document.getElementById('root')}
                    className="modal"
                >
                    <h2>Are u sure want to release {selectedPokemonName}?</h2>
                    <div><img src={selectedPokemonPicture} alt="Pokemon" style={{maxWidth: '100%'}}/></div>
                    <button onClick={handleRelease}>Yes</button>
                    <button onClick={() => setShowSureModal(false)}>cancel</button>
                </Modal>
                <Modal
                    isOpen={showSuccessModal}
                    onRequestClose={() => setShowSuccessModal(false)}
                    contentLabel="Success Modal"
                    appElement={document.getElementById('root')}
                    className="modal"
                >
                    <h2>Success Release {selectedPokemonName}</h2>
                    <div><img src={selectedPokemonPicture} alt="Pokemon" style={{maxWidth: '100%'}}/></div>
                    <button onClick={() => {
                        setShowSuccessModal(false);
                        window.location.reload();
                    }}>Close
                    </button>

                </Modal>
                <Modal
                    isOpen={showFailureModal}
                    onRequestClose={() => setShowFailureModal(false)}
                    contentLabel="Failure Modal"
                    appElement={document.getElementById('root')}
                    className="modal"
                >
                    <h2>Fail Release {selectedPokemonName}</h2>
                    <div><img src={selectedPokemonPicture} alt="Pokemon" style={{maxWidth: '100%'}}/></div>
                    <button onClick={() => setShowFailureModal(false)}>Close</button>
                </Modal>
            </div>
        </div>
    );
};

export default MyPokemon;
