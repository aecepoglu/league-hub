target="$(pwd)/$1"
echo "downloading assets into $target"

echo "graphiql.min.js"
wget --quiet https://github.com/graphql/graphiql/releases/download/v0.11.11/graphiql.min.js $target/
echo "graphiql.js"
wget --quiet https://github.com/graphql/graphiql/releases/download/v0.11.11/graphiql.js $target/
echo "graphiql.css"
wget --quiet https://github.com/graphql/graphiql/releases/download/v0.11.11/graphiql.css $target/
