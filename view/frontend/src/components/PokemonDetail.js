import React, { useEffect, useState } from 'react';
import {Link, useParams} from 'react-router-dom';
import Modal from 'react-modal';
import '../PokemonDetail.css';
import '../MenuBar.css';

const PokemonDetail = () => {
    const { pokemonName } = useParams();
    const [pokemonDetails, setPokemonDetails] = useState(null);
    const [showCatchModal, setShowCatchModal] = useState(false);
    const [showSuccessModal, setShowSuccessModal] = useState(false);
    const [showFailureModal, setShowFailureModal] = useState(false);
    const [showInsertModal, setShowInsertModal] = useState(false);
    const [renameCount, setRenameCount] = useState(-1);
    const [pokeName, setPokeName] = useState(null);

    useEffect(() => {
        fetch(`http://localhost:1323/${pokemonName}`)
            .then((response) => response.json())
            .then((data) => {
                setPokemonDetails(data);
                setPokeName(data.name);
            });
    }, [pokemonName]);

        const handleButtonClick = async () => {
        const response = await fetch('http://localhost:1323/catch');
        const result = await response.json();
        setShowCatchModal(true);
        // Wait for 3 seconds before displaying the success/not success message
        setTimeout(() => {
            setShowCatchModal(false)
            // Show the success or failure modal based on the API response
            if (result) {
                setShowSuccessModal(true);
            } else {
                setShowFailureModal(true);
            }
        }, 3000);
    };

    const handleRenameClick = async () => {
        setPokeName(null)
        const response = await fetch(`http://localhost:1323/rename?name=${pokemonDetails.name}&hit=${renameCount + 1}`);
        const data = await response.json();
        setPokeName(data)
        // Update the rename count and close the modal
        setRenameCount(renameCount + 1);
    };

    const handleInsert = async () => {
        try {
            const response = await fetch('http://localhost:1323/add', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    name: pokeName,
                    picture: pokemonDetails.sprites,
                }),
            });

            if (!response.ok) {
                throw new Error(`Failed to insert data. Status: ${response.status}`);
            }

            setShowSuccessModal(false);
            setShowInsertModal(true)
        } catch (error) {
            // Handle the error
            console.error('Error inserting data:', error.message);
        }
    };

    const handleSuccessModalClose = () => {
        setPokeName(pokemonDetails.name)
        setRenameCount(-1);
        setShowSuccessModal(false)
    };

    const handleInsertModalClose = () => {
        setPokeName(pokemonDetails.name)
        setRenameCount(-1);
        setShowInsertModal(false)
    };

    if (!pokemonDetails) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <div className="menu-bar">
                <Link to="/" className="menu-item">
                    <img src={'/pokeball.png'} alt="Custom Icon"
                         style={{marginRight: '9px', width: '30px', height: '20px'}}/>
                </Link>
                <Link to="/mypokemon" className="menu-item">
                    My Pokemon List
                </Link>
            </div>
            <div className="pokemon-detail-container">
                <h1>{pokemonDetails.name}</h1>
                <div><img src={pokemonDetails.sprites} alt="Pokemon" className="pokemon-image"/></div>
                <div className="catch">
                    <button onClick={handleButtonClick}>Catch Pokemon</button>
                </div>

                <Modal
                    isOpen={showCatchModal}
                    onRequestClose={() => setShowCatchModal(false)}
                    contentLabel="Poke Modal"
                    appElement={document.getElementById('root')}
                    className="modal"
                >
                    <img src={`/catching-pokemon.gif`} alt="GIF" className="modal-image"/>
                </Modal>

                <Modal
                    isOpen={showSuccessModal}
                    onRequestClose={() => setShowSuccessModal(false)}
                    contentLabel="Success Modal"
                    appElement={document.getElementById('root')}
                    className="modal"
                >
                    <img src={pokemonDetails.sprites} alt="Pokemon" style={{maxWidth: '100%'}}/>
                    <h2>You Success Catch!</h2>
                    <h2>{pokeName}</h2>
                    <button onClick={handleRenameClick}>Rename</button>
                    <button onClick={handleInsert}>Insert</button>
                    <button onClick={handleSuccessModalClose}>Close</button>
                </Modal>

                <Modal
                    isOpen={showInsertModal}
                    onRequestClose={() => setShowInsertModal(false)}
                    contentLabel="Insert Modal"
                    appElement={document.getElementById('root')}
                    className="modal"
                >
                    <img src={pokemonDetails.sprites} alt={"Pokemon"} style={{maxWidth: '100%'}}/>
                    <h2>Success insert {pokeName} to your inventory! </h2>
                    <button onClick={handleInsertModalClose}>Close</button>
                </Modal>

                <Modal
                    isOpen={showFailureModal}
                    onRequestClose={() => setShowFailureModal(false)}
                    contentLabel="Failure Modal"
                    appElement={document.getElementById('root')}
                    className="modal"
                >
                    <img src={"/fail.gif"} alt={"Pokemon"} style={{maxWidth: '50%'}}/>
                    <h1>You Fail Catch {pokeName}</h1>
                    <button onClick={() => setShowFailureModal(false)}>Close</button>
                </Modal>

                <div className="section">
                    <h1>Moves</h1>
                    <div className="moves-card">
                        {pokemonDetails.moves.map((move) => (
                            <div key={move} className="move-item">
                                {move}
                            </div>
                        ))}
                    </div>
                </div>
                <div className="section">
                    <h1>Types</h1>
                    <div className="types-card">
                        {pokemonDetails.types.map((type) => (
                            <div key={type} className="type-item">
                                {type}
                            </div>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    );
};

export default PokemonDetail;

