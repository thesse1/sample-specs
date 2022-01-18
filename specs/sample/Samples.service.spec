name: Samples
version: ""
description: |
    Services for handling Sample.
lifecycle: null
__proto:
    package: sample
    targetfile: sampleservice.proto
    imports:
        - google/api/annotations.proto
        - sample/reqmsgs.proto
        - google/protobuf/empty.proto
        - furo/signatures/signatures.proto
        - google/protobuf/field_mask.proto
        - sample/sample.proto
    options:
        go_package: github.com/thesse1/sample-specs/dist/pb/sample;samplepb
        java_multiple_files: "true"
        java_outer_classname: SampleserviceProto
        java_package: com.example.tutorial.sample
services:
    List:
        description: Returns a list of Samples.
        data:
            request: google.protobuf.Empty
            response: sample.SampleCollection
            bodyfield: body
        deeplink:
            description: 'List: GET /sample google.protobuf.Empty , sample.SampleCollection #Returns a list of Samples.'
            href: /sample
            method: GET
            rel: list
        query:
            q:
                constraints: {}
                description: Search query param.
                meta: null
                type: string
            filter:
                constraints: {}
                description: Filter query param
                meta: null
                type: string
            page:
                constraints: {}
                description: Pagination query param.
                meta: null
                type: string
        rpc_name: ListSamples
    Get:
        description: Returns a single Sample.
        data:
            request: google.protobuf.Empty
            response: sample.SampleEntity
            bodyfield: body
        deeplink:
            description: 'Get: GET /sample/{smp} google.protobuf.Empty , sample.SampleEntity #Returns a single Sample.'
            href: /sample/{smp}
            method: GET
            rel: self
        query:
            smp:
                constraints: {}
                description: The query param smp stands for the id of a Sample.
                meta: null
                type: string
        rpc_name: GetSample
    Create:
        description: Creates a Sample.
        data:
            request: sample.Sample
            response: furo.signatures.EmptyEntity
            bodyfield: body
        deeplink:
            description: 'Create: POST /sample sample.Sample , furo.signatures.EmptyEntity #Creates a Sample.'
            href: /sample
            method: POST
            rel: create
        query: {}
        rpc_name: CreateSample
    Update:
        description: Update an existing  Sample. PATCH is also supported
        data:
            request: sample.Sample
            response: sample.SampleEntity
            bodyfield: body
        deeplink:
            description: 'Update: PUT /sample/{smp} sample.Sample , sample.SampleEntity #Update an existing  Sample. PATCH is also supported'
            href: /sample/{smp}
            method: PUT
            rel: update
        query:
            smp:
                constraints: {}
                description: The query param smp stands for the id of a  Sample.
                meta: null
                type: string
            update_mask:
                constraints: {}
                description: Needed to patch a record
                meta: null
                type: google.protobuf.FieldMask
        rpc_name: UpdateSample
    Delete:
        description: Delete a Sample
        data:
            request: google.protobuf.Empty
            response: google.protobuf.Empty
            bodyfield: body
        deeplink:
            description: 'Delete: DELETE /sample/{smp} google.protobuf.Empty , google.protobuf.Empty #Delete a Sample'
            href: /sample/{smp}
            method: DELETE
            rel: delete
        query:
            smp:
                constraints: {}
                description: The query param smp stands for the id of a  Sample.
                meta: null
                type: string
        rpc_name: DeleteSample
