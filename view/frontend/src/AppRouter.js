import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import PokemonList from "./components/PokemonList";
import PokemonDetail from "./components/PokemonDetail";
import MyPokemon from "./components/MyPokemon";

const AppRouter = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<PokemonList />} />
                <Route path="/:pokemonName" element={<PokemonDetail />} />
                <Route path="/mypokemon" element={<MyPokemon />} />
            </Routes>
        </Router>
    );
};

export default AppRouter;
