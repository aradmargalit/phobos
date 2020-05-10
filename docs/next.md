# Next Steps

There is a lot that I wished I'd done differently in this project, so here's a list of next steps once I get time and motivation and tackle these.

## Frontend :computer:

- [ ] Add unit tests, potentially moving from Enzyme to React Testing Library
- [ ] Stop setting component state from API helper functions
- [ ] Create containers to handle state and network calls, leaving "dumb" UI components to act as pure render functions.
- [ ] Responsive UI for mobile users

## Backend :rocket:

- [ ] Unit tests
- [ ] Decouple authentication flows from Gin
- [ ] Use a real logger instead of `fmt.Println`
- [ ] Remove as many `panic`s as possible.
- [x] Hide concrete implementations (database, service) behind interfaces
- [x] Separate transport layer from service layer, and separate service layer from data layer.
