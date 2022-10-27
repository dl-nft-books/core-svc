package helpers

import (
	"bytes"

	chunker "github.com/ipfs/go-ipfs-chunker"
	dagMock "github.com/ipfs/go-merkledag/test"
	"github.com/ipfs/go-unixfs/importer"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func PrecalculateIPFSHash(raw []byte) (string, error) {
	reader := bytes.NewReader(raw)
	dagService := dagMock.Mock()
	chunkSplitter := chunker.DefaultSplitter(reader)

	node, err := importer.BuildDagFromReader(dagService, chunkSplitter)
	if err != nil {
		return "", errors.Wrap(err, "failed to build dag file")
	}

	return node.Cid().Hash().B58String(), nil
}
