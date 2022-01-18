name: Sample
type: Sample
description: A sample type
lifecycle: null
__proto:
    package: sample
    targetfile: sample.proto
    imports: []
    options:
        go_package: github.com/thesse1/sample-specs/dist/pb/sample;samplepb
        java_multiple_files: "true"
        java_outer_classname: SampleProto
        java_package: com.example.tutorial.sample
fields:
    id:
        type: string
        description: ULID as required string.
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: sample.sample.id.placeholder
            hint: ""
            label: sample.sample.id.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints:
            required:
                is: "true"
                message: id is required
    display_name:
        type: string
        description: Readonly display name
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: sample.sample.displayname.placeholder
            hint: ""
            label: sample.sample.displayname.label
            options:
                flags: []
                list: []
            readonly: true
            repeated: false
            typespecific: null
        constraints: {}
    in_edit:
        type: bool
        description: True when a draft exists
        __proto:
            number: 3
        __ui: null
        meta:
            default: ""
            placeholder: sample.sample.inedit.placeholder
            hint: ""
            label: sample.sample.inedit.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    tags:
        type: string
        description: Repeated string
        __proto:
            number: 4
        __ui: null
        meta:
            default: ""
            placeholder: sample.sample.tags.placeholder
            hint: ""
            label: sample.sample.tags.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    this:
        type: string
        description: Oneof example
        __proto:
            number: 5
            oneof: oneofname
        __ui: null
        meta:
            default: ""
            placeholder: sample.sample.this.placeholder
            hint: ""
            label: sample.sample.this.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    is:
        type: string
        description: Oneof example
        __proto:
            number: 6
        __ui: null
        meta:
            default: ""
            placeholder: sample.sample.is.placeholder
            hint: ""
            label: sample.sample.is.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    val:
        type: int32
        description: Oneof example
        __proto:
            number: 7
        __ui: null
        meta:
            default: "42"
            placeholder: sample.sample.val.placeholder
            hint: ""
            label: sample.sample.val.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
