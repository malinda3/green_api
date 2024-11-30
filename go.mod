module project

go 1.23.2

require backend/greenapi v0.0.0
require github.com/rs/cors v1.11.1

replace backend/greenapi => ./backend/greenapi
