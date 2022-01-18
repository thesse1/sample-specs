name: ListSamplesRequest
type: ListSamplesRequest
description: request message for ListSamples
lifecycle: null
__proto:
    package: sample
    targetfile: reqmsgs.proto
    imports:
        - google/protobuf/empty.proto
    options:
        go_package: github.com/thesse1/sample-specs/dist/pb/sample;samplepb
        java_multiple_files: "true"
        java_outer_classname: ReqmsgsProto
        java_package: com.example.tutorial.sample
fields:
    body:
        type: .google.protobuf.Empty
        description: Body with google.protobuf.Empty
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: sample.listsamplesrequest.body.placeholder
            hint: ""
            label: sample.listsamplesrequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    q:
        type: string
        description: Search query param.
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: sample.listsamplesrequest.q.placeholder
            hint: ""
            label: sample.listsamplesrequest.q.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    filter:
        type: string
        description: Filter query param
        __proto:
            number: 3
        __ui: null
        meta:
            default: ""
            placeholder: sample.listsamplesrequest.filter.placeholder
            hint: ""
            label: sample.listsamplesrequest.filter.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    page:
        type: string
        description: Pagination query param.
        __proto:
            number: 4
        __ui: null
        meta:
            default: ""
            placeholder: sample.listsamplesrequest.page.placeholder
            hint: ""
            label: sample.listsamplesrequest.page.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
