# Chess Engine

A basic chess engine I'm slowly working on when I feel up to it. Nothing revolutionary, but a fun and challenging problem.

Goals:

- [x] Read and output board states (FEN)
- [ ] List all valid moves for a given position
- [ ] Make a move, updating the board state
- [ ] Validate listing of moves by comparing total possible states 6 moves deep vs stockfish
- [ ] Create a UI you can play chess on using `.wasm` binaries to run the calculations
- [ ] Optimize!

## Project Structure

`engine/`
   
    The golang project that does all the heavy lifting

`ui/`

    The web UI for playing chess against the engine
