name: CreateSampleRequest
type: CreateSampleRequest
description: request message for CreateSample
lifecycle: null
__proto:
    package: sample
    targetfile: reqmsgs.proto
    imports:
        - sample/sample.proto
    options:
        go_package: github.com/thesse1/sample-specs/dist/pb/sample;samplepb
        java_multiple_files: "true"
        java_outer_classname: ReqmsgsProto
        java_package: com.example.tutorial.sample
fields:
    body:
        type: .sample.Sample
        description: Body with sample.Sample
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: sample.createsamplerequest.body.placeholder
            hint: ""
            label: sample.createsamplerequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
