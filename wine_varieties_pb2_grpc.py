# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import wine_varieties_pb2 as wine__varieties__pb2


class WineClassifierServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetWineVariety = channel.unary_unary(
                '/pb.WineClassifierService/GetWineVariety',
                request_serializer=wine__varieties__pb2.WineReviewRequest.SerializeToString,
                response_deserializer=wine__varieties__pb2.WineReviewResponse.FromString,
                )


class WineClassifierServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetWineVariety(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_WineClassifierServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetWineVariety': grpc.unary_unary_rpc_method_handler(
                    servicer.GetWineVariety,
                    request_deserializer=wine__varieties__pb2.WineReviewRequest.FromString,
                    response_serializer=wine__varieties__pb2.WineReviewResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'pb.WineClassifierService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class WineClassifierService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetWineVariety(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/pb.WineClassifierService/GetWineVariety',
            wine__varieties__pb2.WineReviewRequest.SerializeToString,
            wine__varieties__pb2.WineReviewResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
