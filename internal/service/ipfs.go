package service

import (
	"context"
	"fmt"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs/core/coreapi"
	//"github.com/ipfs/go-ipfs/core/coreunix"
	//"github.com/ipfs/go-ipfs/core/node"
	//"github.com/ipfs/interface-go-ipfs-core/path"
	//"path/filepath"

	core "github.com/ipfs/go-ipfs/core"
	"github.com/sirupsen/logrus"
	"os"
)

type IpfsService interface {
	SaveFile(context.Context, *os.File) (string, error)
}

type ipfsService struct {
	node *core.IpfsNode
}

func NewIPFSService() (IpfsService, error) {
	node, err := core.NewNode(context.Background(), &core.BuildCfg{
		Online: true,
	})

	if err != nil {
		logrus.Error("error while initializing  IPFS node. Error: ", err.Error())
		return nil, err
	}

	return &ipfsService{
		node: node,
	}, nil
}

func (i ipfsService) SaveFile(ctx context.Context, file *os.File) (string, error) {
	if file == nil {
		return "", fmt.Errorf("file not found")
	}

	coreAPI, err := coreapi.NewCoreAPI(i.node)
	if err != nil {
		return "", err
	}

	//stat, err := coreAPI.Block().Put(ctx, file)
	//if err != nil {
	//	fmt.Println("error occurred here-->", err.Error())
	//	return "", err
	//}
	//
	//return stat.Path().Cid().String(), nil
	//

	//adder, err := coreunix.NewAdder(ctx, i.node.Pinning, i.node.Blockstore, i.node.DAG)
	//if err != nil {
	//	return "", err
	//}
	//
	//fileReader := files.NewReaderFile(file)
	//node, err := adder.AddAllAndPin(fileReader)
	//if err != nil {
	//	return "", err
	//}
	//
	//stat, err := node.Stat()
	//if err != nil {
	//	return "", err
	//}
	//return stat.Hash, nil

	path, err := coreAPI.Unixfs().Add(ctx, files.NewReaderFile(file))
	if err != nil {
		logrus.Error("error while adding file. Error: ", err.Error())
		return "", err
	}

	return path.Cid().String(), nil

}
