goctl model mysql datasource \
    --url 'root:root542@tcp(127.0.0.1:3307)/qq_forum' \
    -d ./model/user \
    -c \
    -t user \
    --style goZero

goctl model mysql datasource \
    --url 'root:root542@tcp(127.0.0.1:3307)/qq_forum' \
    -d ./model/post \
    -c \
    -t post \
    --style goZero

goctl model mysql datasource \
    --url 'root:root542@tcp(127.0.0.1:3307)/qq_forum' \
    -d ./model/comment \
    -c \
    -t comment \
    --style goZero