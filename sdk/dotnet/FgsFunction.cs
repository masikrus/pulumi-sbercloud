// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sbercloud
{
    [SbercloudResourceType("sbercloud:index/fgsFunction:FgsFunction")]
    public partial class FgsFunction : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The agency configuration of the function.
        /// </summary>
        [Output("agency")]
        public Output<string?> Agency { get; private set; } = null!;

        /// <summary>
        /// The group to which the function belongs.
        /// </summary>
        [Output("app")]
        public Output<string?> App { get; private set; } = null!;

        /// <summary>
        /// The execution agency enables you to obtain a token or an AK/SK for accessing other cloud services.
        /// </summary>
        [Output("appAgency")]
        public Output<string> AppAgency { get; private set; } = null!;

        /// <summary>
        /// The name of the function file.
        /// </summary>
        [Output("codeFilename")]
        public Output<string> CodeFilename { get; private set; } = null!;

        /// <summary>
        /// The code type of the function.
        /// </summary>
        [Output("codeType")]
        public Output<string> CodeType { get; private set; } = null!;

        /// <summary>
        /// The URL where the function code is stored in OBS.
        /// </summary>
        [Output("codeUrl")]
        public Output<string?> CodeUrl { get; private set; } = null!;

        /// <summary>
        /// The number of concurrent requests of the function.
        /// </summary>
        [Output("concurrencyNum")]
        public Output<int> ConcurrencyNum { get; private set; } = null!;

        /// <summary>
        /// The custom image configuration of the function.
        /// </summary>
        [Output("customImage")]
        public Output<Outputs.FgsFunctionCustomImage> CustomImage { get; private set; } = null!;

        /// <summary>
        /// The ID list of the dependencies.
        /// </summary>
        [Output("dependLists")]
        public Output<ImmutableArray<string>> DependLists { get; private set; } = null!;

        /// <summary>
        /// The description of the function.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The private DNS configuration of the function network.
        /// </summary>
        [Output("dnsList")]
        public Output<string> DnsList { get; private set; } = null!;

        /// <summary>
        /// Whether the authentication in the request header is enabled.
        /// </summary>
        [Output("enableAuthInHeader")]
        public Output<bool?> EnableAuthInHeader { get; private set; } = null!;

        /// <summary>
        /// Whether the class isolation is enabled for the JAVA runtime functions.
        /// </summary>
        [Output("enableClassIsolation")]
        public Output<bool?> EnableClassIsolation { get; private set; } = null!;

        /// <summary>
        /// Whether the dynamic memory configuration is enabled.
        /// </summary>
        [Output("enableDynamicMemory")]
        public Output<bool?> EnableDynamicMemory { get; private set; } = null!;

        /// <summary>
        /// The key/value information defined to be encrypted for the function.
        /// </summary>
        [Output("encryptedUserData")]
        public Output<string?> EncryptedUserData { get; private set; } = null!;

        /// <summary>
        /// The ID of the enterprise project to which the function belongs.
        /// </summary>
        [Output("enterpriseProjectId")]
        public Output<string> EnterpriseProjectId { get; private set; } = null!;

        /// <summary>
        /// The size of the function ephemeral storage.
        /// </summary>
        [Output("ephemeralStorage")]
        public Output<int> EphemeralStorage { get; private set; } = null!;

        /// <summary>
        /// The function code.
        /// </summary>
        [Output("funcCode")]
        public Output<string?> FuncCode { get; private set; } = null!;

        /// <summary>
        /// The list of function mount configuration.
        /// </summary>
        [Output("funcMounts")]
        public Output<ImmutableArray<Outputs.FgsFunctionFuncMount>> FuncMounts { get; private set; } = null!;

        /// <summary>
        /// The description of the function.
        /// </summary>
        [Output("functiongraphVersion")]
        public Output<string> FunctiongraphVersion { get; private set; } = null!;

        /// <summary>
        /// The GPU memory size allocated to the function, in MByte (MB).
        /// </summary>
        [Output("gpuMemory")]
        public Output<int?> GpuMemory { get; private set; } = null!;

        /// <summary>
        /// The GPU type of the function.
        /// </summary>
        [Output("gpuType")]
        public Output<string?> GpuType { get; private set; } = null!;

        /// <summary>
        /// The entry point of the function.
        /// </summary>
        [Output("handler")]
        public Output<string> Handler { get; private set; } = null!;

        /// <summary>
        /// The heartbeat handler of the function.
        /// </summary>
        [Output("heartbeatHandler")]
        public Output<string?> HeartbeatHandler { get; private set; } = null!;

        /// <summary>
        /// The initializer of the function.
        /// </summary>
        [Output("initializerHandler")]
        public Output<string> InitializerHandler { get; private set; } = null!;

        /// <summary>
        /// The maximum duration the function can be initialized.
        /// </summary>
        [Output("initializerTimeout")]
        public Output<int> InitializerTimeout { get; private set; } = null!;

        /// <summary>
        /// Whether the function is a stateful function.
        /// </summary>
        [Output("isStatefulFunction")]
        public Output<bool?> IsStatefulFunction { get; private set; } = null!;

        /// <summary>
        /// The LTS group ID for collecting logs.
        /// </summary>
        [Output("logGroupId")]
        public Output<string> LogGroupId { get; private set; } = null!;

        /// <summary>
        /// The LTS group name for collecting logs.
        /// </summary>
        [Output("logGroupName")]
        public Output<string> LogGroupName { get; private set; } = null!;

        /// <summary>
        /// The LTS stream ID for collecting logs.
        /// </summary>
        [Output("logStreamId")]
        public Output<string> LogStreamId { get; private set; } = null!;

        /// <summary>
        /// The LTS stream name for collecting logs.
        /// </summary>
        [Output("logStreamName")]
        public Output<string> LogStreamName { get; private set; } = null!;

        /// <summary>
        /// The maximum number of instances of the function.
        /// </summary>
        [Output("maxInstanceNum")]
        public Output<string> MaxInstanceNum { get; private set; } = null!;

        /// <summary>
        /// The memory size allocated to the function, in MByte (MB).
        /// </summary>
        [Output("memorySize")]
        public Output<int> MemorySize { get; private set; } = null!;

        /// <summary>
        /// The mount user group ID.
        /// </summary>
        [Output("mountUserGroupId")]
        public Output<int> MountUserGroupId { get; private set; } = null!;

        /// <summary>
        /// The mount user ID.
        /// </summary>
        [Output("mountUserId")]
        public Output<int> MountUserId { get; private set; } = null!;

        /// <summary>
        /// The name of the function.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network configuration of the function.
        /// </summary>
        [Output("networkController")]
        public Output<Outputs.FgsFunctionNetworkController?> NetworkController { get; private set; } = null!;

        /// <summary>
        /// The network ID of subnet.
        /// </summary>
        [Output("networkId")]
        public Output<string?> NetworkId { get; private set; } = null!;

        [Output("package")]
        public Output<string?> Package { get; private set; } = null!;

        /// <summary>
        /// The VPC CIDR blocks used in the function code to detect whether it conflicts with the VPC CIDR blocks used by the
        /// service.
        /// </summary>
        [Output("peeringCidr")]
        public Output<string?> PeeringCidr { get; private set; } = null!;

        /// <summary>
        /// The pre-stop handler of a function.
        /// </summary>
        [Output("preStopHandler")]
        public Output<string?> PreStopHandler { get; private set; } = null!;

        /// <summary>
        /// The maximum duration that the function can be initialized.
        /// </summary>
        [Output("preStopTimeout")]
        public Output<int?> PreStopTimeout { get; private set; } = null!;

        /// <summary>
        /// The region where the function is located.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The reserved instance policies of the function.
        /// </summary>
        [Output("reservedInstances")]
        public Output<ImmutableArray<Outputs.FgsFunctionReservedInstance>> ReservedInstances { get; private set; } = null!;

        /// <summary>
        /// The restore hook handler of the function.
        /// </summary>
        [Output("restoreHookHandler")]
        public Output<string?> RestoreHookHandler { get; private set; } = null!;

        /// <summary>
        /// The timeout of the function restore hook.
        /// </summary>
        [Output("restoreHookTimeout")]
        public Output<int?> RestoreHookTimeout { get; private set; } = null!;

        /// <summary>
        /// The environment for executing the function.
        /// </summary>
        [Output("runtime")]
        public Output<string> Runtime { get; private set; } = null!;

        /// <summary>
        /// The key/value pairs to associate with the function.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The timeout interval of the function, in seconds.
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// The URN (Uniform Resource Name) of the function.
        /// </summary>
        [Output("urn")]
        public Output<string> Urn { get; private set; } = null!;

        /// <summary>
        /// The key/value information defined for the function.
        /// </summary>
        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

        /// <summary>
        /// The version of the function.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The versions management of the function.
        /// </summary>
        [Output("versions")]
        public Output<ImmutableArray<Outputs.FgsFunctionVersion>> Versions { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC to which the function belongs.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;

        [Output("xrole")]
        public Output<string?> Xrole { get; private set; } = null!;


        /// <summary>
        /// Create a FgsFunction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FgsFunction(string name, FgsFunctionArgs args, CustomResourceOptions? options = null)
            : base("sbercloud:index/fgsFunction:FgsFunction", name, args ?? new FgsFunctionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FgsFunction(string name, Input<string> id, FgsFunctionState? state = null, CustomResourceOptions? options = null)
            : base("sbercloud:index/fgsFunction:FgsFunction", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "encryptedUserData",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FgsFunction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FgsFunction Get(string name, Input<string> id, FgsFunctionState? state = null, CustomResourceOptions? options = null)
        {
            return new FgsFunction(name, id, state, options);
        }
    }

    public sealed class FgsFunctionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The agency configuration of the function.
        /// </summary>
        [Input("agency")]
        public Input<string>? Agency { get; set; }

        /// <summary>
        /// The group to which the function belongs.
        /// </summary>
        [Input("app")]
        public Input<string>? App { get; set; }

        /// <summary>
        /// The execution agency enables you to obtain a token or an AK/SK for accessing other cloud services.
        /// </summary>
        [Input("appAgency")]
        public Input<string>? AppAgency { get; set; }

        /// <summary>
        /// The name of the function file.
        /// </summary>
        [Input("codeFilename")]
        public Input<string>? CodeFilename { get; set; }

        /// <summary>
        /// The code type of the function.
        /// </summary>
        [Input("codeType")]
        public Input<string>? CodeType { get; set; }

        /// <summary>
        /// The URL where the function code is stored in OBS.
        /// </summary>
        [Input("codeUrl")]
        public Input<string>? CodeUrl { get; set; }

        /// <summary>
        /// The number of concurrent requests of the function.
        /// </summary>
        [Input("concurrencyNum")]
        public Input<int>? ConcurrencyNum { get; set; }

        /// <summary>
        /// The custom image configuration of the function.
        /// </summary>
        [Input("customImage")]
        public Input<Inputs.FgsFunctionCustomImageArgs>? CustomImage { get; set; }

        [Input("dependLists")]
        private InputList<string>? _dependLists;

        /// <summary>
        /// The ID list of the dependencies.
        /// </summary>
        public InputList<string> DependLists
        {
            get => _dependLists ?? (_dependLists = new InputList<string>());
            set => _dependLists = value;
        }

        /// <summary>
        /// The description of the function.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The private DNS configuration of the function network.
        /// </summary>
        [Input("dnsList")]
        public Input<string>? DnsList { get; set; }

        /// <summary>
        /// Whether the authentication in the request header is enabled.
        /// </summary>
        [Input("enableAuthInHeader")]
        public Input<bool>? EnableAuthInHeader { get; set; }

        /// <summary>
        /// Whether the class isolation is enabled for the JAVA runtime functions.
        /// </summary>
        [Input("enableClassIsolation")]
        public Input<bool>? EnableClassIsolation { get; set; }

        /// <summary>
        /// Whether the dynamic memory configuration is enabled.
        /// </summary>
        [Input("enableDynamicMemory")]
        public Input<bool>? EnableDynamicMemory { get; set; }

        [Input("encryptedUserData")]
        private Input<string>? _encryptedUserData;

        /// <summary>
        /// The key/value information defined to be encrypted for the function.
        /// </summary>
        public Input<string>? EncryptedUserData
        {
            get => _encryptedUserData;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _encryptedUserData = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The ID of the enterprise project to which the function belongs.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        /// <summary>
        /// The size of the function ephemeral storage.
        /// </summary>
        [Input("ephemeralStorage")]
        public Input<int>? EphemeralStorage { get; set; }

        /// <summary>
        /// The function code.
        /// </summary>
        [Input("funcCode")]
        public Input<string>? FuncCode { get; set; }

        [Input("funcMounts")]
        private InputList<Inputs.FgsFunctionFuncMountArgs>? _funcMounts;

        /// <summary>
        /// The list of function mount configuration.
        /// </summary>
        public InputList<Inputs.FgsFunctionFuncMountArgs> FuncMounts
        {
            get => _funcMounts ?? (_funcMounts = new InputList<Inputs.FgsFunctionFuncMountArgs>());
            set => _funcMounts = value;
        }

        /// <summary>
        /// The description of the function.
        /// </summary>
        [Input("functiongraphVersion")]
        public Input<string>? FunctiongraphVersion { get; set; }

        /// <summary>
        /// The GPU memory size allocated to the function, in MByte (MB).
        /// </summary>
        [Input("gpuMemory")]
        public Input<int>? GpuMemory { get; set; }

        /// <summary>
        /// The GPU type of the function.
        /// </summary>
        [Input("gpuType")]
        public Input<string>? GpuType { get; set; }

        /// <summary>
        /// The entry point of the function.
        /// </summary>
        [Input("handler")]
        public Input<string>? Handler { get; set; }

        /// <summary>
        /// The heartbeat handler of the function.
        /// </summary>
        [Input("heartbeatHandler")]
        public Input<string>? HeartbeatHandler { get; set; }

        /// <summary>
        /// The initializer of the function.
        /// </summary>
        [Input("initializerHandler")]
        public Input<string>? InitializerHandler { get; set; }

        /// <summary>
        /// The maximum duration the function can be initialized.
        /// </summary>
        [Input("initializerTimeout")]
        public Input<int>? InitializerTimeout { get; set; }

        /// <summary>
        /// Whether the function is a stateful function.
        /// </summary>
        [Input("isStatefulFunction")]
        public Input<bool>? IsStatefulFunction { get; set; }

        /// <summary>
        /// The LTS group ID for collecting logs.
        /// </summary>
        [Input("logGroupId")]
        public Input<string>? LogGroupId { get; set; }

        /// <summary>
        /// The LTS group name for collecting logs.
        /// </summary>
        [Input("logGroupName")]
        public Input<string>? LogGroupName { get; set; }

        /// <summary>
        /// The LTS stream ID for collecting logs.
        /// </summary>
        [Input("logStreamId")]
        public Input<string>? LogStreamId { get; set; }

        /// <summary>
        /// The LTS stream name for collecting logs.
        /// </summary>
        [Input("logStreamName")]
        public Input<string>? LogStreamName { get; set; }

        /// <summary>
        /// The maximum number of instances of the function.
        /// </summary>
        [Input("maxInstanceNum")]
        public Input<string>? MaxInstanceNum { get; set; }

        /// <summary>
        /// The memory size allocated to the function, in MByte (MB).
        /// </summary>
        [Input("memorySize", required: true)]
        public Input<int> MemorySize { get; set; } = null!;

        /// <summary>
        /// The mount user group ID.
        /// </summary>
        [Input("mountUserGroupId")]
        public Input<int>? MountUserGroupId { get; set; }

        /// <summary>
        /// The mount user ID.
        /// </summary>
        [Input("mountUserId")]
        public Input<int>? MountUserId { get; set; }

        /// <summary>
        /// The name of the function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network configuration of the function.
        /// </summary>
        [Input("networkController")]
        public Input<Inputs.FgsFunctionNetworkControllerArgs>? NetworkController { get; set; }

        /// <summary>
        /// The network ID of subnet.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("package")]
        public Input<string>? Package { get; set; }

        /// <summary>
        /// The VPC CIDR blocks used in the function code to detect whether it conflicts with the VPC CIDR blocks used by the
        /// service.
        /// </summary>
        [Input("peeringCidr")]
        public Input<string>? PeeringCidr { get; set; }

        /// <summary>
        /// The pre-stop handler of a function.
        /// </summary>
        [Input("preStopHandler")]
        public Input<string>? PreStopHandler { get; set; }

        /// <summary>
        /// The maximum duration that the function can be initialized.
        /// </summary>
        [Input("preStopTimeout")]
        public Input<int>? PreStopTimeout { get; set; }

        /// <summary>
        /// The region where the function is located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("reservedInstances")]
        private InputList<Inputs.FgsFunctionReservedInstanceArgs>? _reservedInstances;

        /// <summary>
        /// The reserved instance policies of the function.
        /// </summary>
        public InputList<Inputs.FgsFunctionReservedInstanceArgs> ReservedInstances
        {
            get => _reservedInstances ?? (_reservedInstances = new InputList<Inputs.FgsFunctionReservedInstanceArgs>());
            set => _reservedInstances = value;
        }

        /// <summary>
        /// The restore hook handler of the function.
        /// </summary>
        [Input("restoreHookHandler")]
        public Input<string>? RestoreHookHandler { get; set; }

        /// <summary>
        /// The timeout of the function restore hook.
        /// </summary>
        [Input("restoreHookTimeout")]
        public Input<int>? RestoreHookTimeout { get; set; }

        /// <summary>
        /// The environment for executing the function.
        /// </summary>
        [Input("runtime", required: true)]
        public Input<string> Runtime { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The key/value pairs to associate with the function.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The timeout interval of the function, in seconds.
        /// </summary>
        [Input("timeout", required: true)]
        public Input<int> Timeout { get; set; } = null!;

        /// <summary>
        /// The key/value information defined for the function.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        [Input("versions")]
        private InputList<Inputs.FgsFunctionVersionArgs>? _versions;

        /// <summary>
        /// The versions management of the function.
        /// </summary>
        public InputList<Inputs.FgsFunctionVersionArgs> Versions
        {
            get => _versions ?? (_versions = new InputList<Inputs.FgsFunctionVersionArgs>());
            set => _versions = value;
        }

        /// <summary>
        /// The ID of the VPC to which the function belongs.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        [Input("xrole")]
        public Input<string>? Xrole { get; set; }

        public FgsFunctionArgs()
        {
        }
        public static new FgsFunctionArgs Empty => new FgsFunctionArgs();
    }

    public sealed class FgsFunctionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The agency configuration of the function.
        /// </summary>
        [Input("agency")]
        public Input<string>? Agency { get; set; }

        /// <summary>
        /// The group to which the function belongs.
        /// </summary>
        [Input("app")]
        public Input<string>? App { get; set; }

        /// <summary>
        /// The execution agency enables you to obtain a token or an AK/SK for accessing other cloud services.
        /// </summary>
        [Input("appAgency")]
        public Input<string>? AppAgency { get; set; }

        /// <summary>
        /// The name of the function file.
        /// </summary>
        [Input("codeFilename")]
        public Input<string>? CodeFilename { get; set; }

        /// <summary>
        /// The code type of the function.
        /// </summary>
        [Input("codeType")]
        public Input<string>? CodeType { get; set; }

        /// <summary>
        /// The URL where the function code is stored in OBS.
        /// </summary>
        [Input("codeUrl")]
        public Input<string>? CodeUrl { get; set; }

        /// <summary>
        /// The number of concurrent requests of the function.
        /// </summary>
        [Input("concurrencyNum")]
        public Input<int>? ConcurrencyNum { get; set; }

        /// <summary>
        /// The custom image configuration of the function.
        /// </summary>
        [Input("customImage")]
        public Input<Inputs.FgsFunctionCustomImageGetArgs>? CustomImage { get; set; }

        [Input("dependLists")]
        private InputList<string>? _dependLists;

        /// <summary>
        /// The ID list of the dependencies.
        /// </summary>
        public InputList<string> DependLists
        {
            get => _dependLists ?? (_dependLists = new InputList<string>());
            set => _dependLists = value;
        }

        /// <summary>
        /// The description of the function.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The private DNS configuration of the function network.
        /// </summary>
        [Input("dnsList")]
        public Input<string>? DnsList { get; set; }

        /// <summary>
        /// Whether the authentication in the request header is enabled.
        /// </summary>
        [Input("enableAuthInHeader")]
        public Input<bool>? EnableAuthInHeader { get; set; }

        /// <summary>
        /// Whether the class isolation is enabled for the JAVA runtime functions.
        /// </summary>
        [Input("enableClassIsolation")]
        public Input<bool>? EnableClassIsolation { get; set; }

        /// <summary>
        /// Whether the dynamic memory configuration is enabled.
        /// </summary>
        [Input("enableDynamicMemory")]
        public Input<bool>? EnableDynamicMemory { get; set; }

        [Input("encryptedUserData")]
        private Input<string>? _encryptedUserData;

        /// <summary>
        /// The key/value information defined to be encrypted for the function.
        /// </summary>
        public Input<string>? EncryptedUserData
        {
            get => _encryptedUserData;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _encryptedUserData = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The ID of the enterprise project to which the function belongs.
        /// </summary>
        [Input("enterpriseProjectId")]
        public Input<string>? EnterpriseProjectId { get; set; }

        /// <summary>
        /// The size of the function ephemeral storage.
        /// </summary>
        [Input("ephemeralStorage")]
        public Input<int>? EphemeralStorage { get; set; }

        /// <summary>
        /// The function code.
        /// </summary>
        [Input("funcCode")]
        public Input<string>? FuncCode { get; set; }

        [Input("funcMounts")]
        private InputList<Inputs.FgsFunctionFuncMountGetArgs>? _funcMounts;

        /// <summary>
        /// The list of function mount configuration.
        /// </summary>
        public InputList<Inputs.FgsFunctionFuncMountGetArgs> FuncMounts
        {
            get => _funcMounts ?? (_funcMounts = new InputList<Inputs.FgsFunctionFuncMountGetArgs>());
            set => _funcMounts = value;
        }

        /// <summary>
        /// The description of the function.
        /// </summary>
        [Input("functiongraphVersion")]
        public Input<string>? FunctiongraphVersion { get; set; }

        /// <summary>
        /// The GPU memory size allocated to the function, in MByte (MB).
        /// </summary>
        [Input("gpuMemory")]
        public Input<int>? GpuMemory { get; set; }

        /// <summary>
        /// The GPU type of the function.
        /// </summary>
        [Input("gpuType")]
        public Input<string>? GpuType { get; set; }

        /// <summary>
        /// The entry point of the function.
        /// </summary>
        [Input("handler")]
        public Input<string>? Handler { get; set; }

        /// <summary>
        /// The heartbeat handler of the function.
        /// </summary>
        [Input("heartbeatHandler")]
        public Input<string>? HeartbeatHandler { get; set; }

        /// <summary>
        /// The initializer of the function.
        /// </summary>
        [Input("initializerHandler")]
        public Input<string>? InitializerHandler { get; set; }

        /// <summary>
        /// The maximum duration the function can be initialized.
        /// </summary>
        [Input("initializerTimeout")]
        public Input<int>? InitializerTimeout { get; set; }

        /// <summary>
        /// Whether the function is a stateful function.
        /// </summary>
        [Input("isStatefulFunction")]
        public Input<bool>? IsStatefulFunction { get; set; }

        /// <summary>
        /// The LTS group ID for collecting logs.
        /// </summary>
        [Input("logGroupId")]
        public Input<string>? LogGroupId { get; set; }

        /// <summary>
        /// The LTS group name for collecting logs.
        /// </summary>
        [Input("logGroupName")]
        public Input<string>? LogGroupName { get; set; }

        /// <summary>
        /// The LTS stream ID for collecting logs.
        /// </summary>
        [Input("logStreamId")]
        public Input<string>? LogStreamId { get; set; }

        /// <summary>
        /// The LTS stream name for collecting logs.
        /// </summary>
        [Input("logStreamName")]
        public Input<string>? LogStreamName { get; set; }

        /// <summary>
        /// The maximum number of instances of the function.
        /// </summary>
        [Input("maxInstanceNum")]
        public Input<string>? MaxInstanceNum { get; set; }

        /// <summary>
        /// The memory size allocated to the function, in MByte (MB).
        /// </summary>
        [Input("memorySize")]
        public Input<int>? MemorySize { get; set; }

        /// <summary>
        /// The mount user group ID.
        /// </summary>
        [Input("mountUserGroupId")]
        public Input<int>? MountUserGroupId { get; set; }

        /// <summary>
        /// The mount user ID.
        /// </summary>
        [Input("mountUserId")]
        public Input<int>? MountUserId { get; set; }

        /// <summary>
        /// The name of the function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network configuration of the function.
        /// </summary>
        [Input("networkController")]
        public Input<Inputs.FgsFunctionNetworkControllerGetArgs>? NetworkController { get; set; }

        /// <summary>
        /// The network ID of subnet.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        [Input("package")]
        public Input<string>? Package { get; set; }

        /// <summary>
        /// The VPC CIDR blocks used in the function code to detect whether it conflicts with the VPC CIDR blocks used by the
        /// service.
        /// </summary>
        [Input("peeringCidr")]
        public Input<string>? PeeringCidr { get; set; }

        /// <summary>
        /// The pre-stop handler of a function.
        /// </summary>
        [Input("preStopHandler")]
        public Input<string>? PreStopHandler { get; set; }

        /// <summary>
        /// The maximum duration that the function can be initialized.
        /// </summary>
        [Input("preStopTimeout")]
        public Input<int>? PreStopTimeout { get; set; }

        /// <summary>
        /// The region where the function is located.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("reservedInstances")]
        private InputList<Inputs.FgsFunctionReservedInstanceGetArgs>? _reservedInstances;

        /// <summary>
        /// The reserved instance policies of the function.
        /// </summary>
        public InputList<Inputs.FgsFunctionReservedInstanceGetArgs> ReservedInstances
        {
            get => _reservedInstances ?? (_reservedInstances = new InputList<Inputs.FgsFunctionReservedInstanceGetArgs>());
            set => _reservedInstances = value;
        }

        /// <summary>
        /// The restore hook handler of the function.
        /// </summary>
        [Input("restoreHookHandler")]
        public Input<string>? RestoreHookHandler { get; set; }

        /// <summary>
        /// The timeout of the function restore hook.
        /// </summary>
        [Input("restoreHookTimeout")]
        public Input<int>? RestoreHookTimeout { get; set; }

        /// <summary>
        /// The environment for executing the function.
        /// </summary>
        [Input("runtime")]
        public Input<string>? Runtime { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The key/value pairs to associate with the function.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The timeout interval of the function, in seconds.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The URN (Uniform Resource Name) of the function.
        /// </summary>
        [Input("urn")]
        public Input<string>? Urn { get; set; }

        /// <summary>
        /// The key/value information defined for the function.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        /// <summary>
        /// The version of the function.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        [Input("versions")]
        private InputList<Inputs.FgsFunctionVersionGetArgs>? _versions;

        /// <summary>
        /// The versions management of the function.
        /// </summary>
        public InputList<Inputs.FgsFunctionVersionGetArgs> Versions
        {
            get => _versions ?? (_versions = new InputList<Inputs.FgsFunctionVersionGetArgs>());
            set => _versions = value;
        }

        /// <summary>
        /// The ID of the VPC to which the function belongs.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        [Input("xrole")]
        public Input<string>? Xrole { get; set; }

        public FgsFunctionState()
        {
        }
        public static new FgsFunctionState Empty => new FgsFunctionState();
    }
}
