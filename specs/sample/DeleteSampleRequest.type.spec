name: DeleteSampleRequest
type: DeleteSampleRequest
description: request message for DeleteSample
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
            placeholder: sample.deletesamplerequest.body.placeholder
            hint: ""
            label: sample.deletesamplerequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    smp:
        type: string
        description: The query param smp stands for the id of a  Sample.
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: sample.deletesamplerequest.smp.placeholder
            hint: ""
            label: sample.deletesamplerequest.smp.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
