package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/qubetics/qubetics-blockchain/v2/types"
	v1 "github.com/qubetics/qubetics-blockchain/v2/x/lease/types/v1"
)

const (
	// gRPC methods for querying lease information
	methodQueryLease             = "/qubetics.lease.v1.QueryService/QueryLease"             // Retrieve a specific lease by ID
	methodQueryLeases            = "/qubetics.lease.v1.QueryService/QueryLeases"            // List leases with pagination
	methodQueryLeasesForNode     = "/qubetics.lease.v1.QueryService/QueryLeasesForNode"     // List leases associated with a specific node
	methodQueryLeasesForProvider = "/qubetics.lease.v1.QueryService/QueryLeasesForProvider" // List leases associated with a specific provider
)

// Lease retrieves details of a specific lease by ID.
// Returns the lease information and any error encountered.
func (c *Client) Lease(ctx context.Context, id uint64) (res *v1.Lease, err error) {
	var (
		resp v1.QueryLeaseResponse
		req  = &v1.QueryLeaseRequest{Id: id}
	)

	// Perform the gRPC query to fetch the lease details.
	if err := c.QueryGRPC(ctx, methodQueryLease, req, &resp); err != nil {
		return nil, IsCodeNotFound(err)
	}

	return &resp.Lease, nil
}

// Leases retrieves a paginated list of all leases.
// Returns the leases, pagination details, and any error encountered.
func (c *Client) Leases(ctx context.Context, pageReq *query.PageRequest) (res []v1.Lease, pageRes *query.PageResponse, err error) {
	var (
		resp v1.QueryLeasesResponse
		req  = &v1.QueryLeasesRequest{Pagination: pageReq}
	)

	// Perform the gRPC query to fetch the leases.
	if err := c.QueryGRPC(ctx, methodQueryLeases, req, &resp); err != nil {
		return nil, nil, err
	}

	return resp.Leases, resp.Pagination, nil
}

// LeasesForNode retrieves leases associated with a specific node address.
// Returns the leases, pagination details, and any error encountered.
func (c *Client) LeasesForNode(ctx context.Context, nodeAddr types.NodeAddress, pageReq *query.PageRequest) (res []v1.Lease, pageRes *query.PageResponse, err error) {
	var (
		resp v1.QueryLeasesForNodeResponse
		req  = &v1.QueryLeasesForNodeRequest{
			Address:    nodeAddr.String(),
			Pagination: pageReq,
		}
	)

	// Perform the gRPC query to fetch leases for the given node.
	if err := c.QueryGRPC(ctx, methodQueryLeasesForNode, req, &resp); err != nil {
		return nil, nil, err
	}

	return resp.Leases, resp.Pagination, nil
}

// LeasesForProvider retrieves leases associated with a specific provider address.
// Returns the leases, pagination details, and any error encountered.
func (c *Client) LeasesForProvider(ctx context.Context, provAddr types.ProvAddress, pageReq *query.PageRequest) (res []v1.Lease, pageRes *query.PageResponse, err error) {
	var (
		resp v1.QueryLeasesForProviderResponse
		req  = &v1.QueryLeasesForProviderRequest{
			Address:    provAddr.String(),
			Pagination: pageReq,
		}
	)

	// Perform the gRPC query to fetch leases for the given provider.
	if err := c.QueryGRPC(ctx, methodQueryLeasesForProvider, req, &resp); err != nil {
		return nil, nil, err
	}

	return resp.Leases, resp.Pagination, nil
}
