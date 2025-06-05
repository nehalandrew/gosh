go build
mkdir -p pkg/usr/local/bin
mv gosh pkg/usr/local/bin/gosh
dpkg-deb --build pkg
lintian pkg.deb