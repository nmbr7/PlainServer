package types

type Ec2_getSubnetsFilter struct {
	/*
	   Name of the field to filter by, as defined by
	   [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSubnets.html).
	   For example, if matching against tag `Name`, use:

	   ```typescript
	   import - as pulumi from "@pulumi/pulumi";
	   import - as aws from "@pulumi/aws";

	   const selected = aws.ec2.getSubnets({
	       filters: [{
	           name: "tag:Name",
	           values: [""],
	       }],
	   });
	   ```
	   ```python
	   import pulumi
	   import pulumi_aws as aws

	   selected = aws.ec2.get_subnets(filters=[aws.ec2.GetSubnetsFilterArgs(
	       name="tag:Name",
	       values=[""],
	   )])
	   ```
	   ```csharp
	   using System.Collections.Generic;
	   using System.Linq;
	   using Pulumi;
	   using Aws = Pulumi.Aws;

	   return await Deployment.RunAsync(() =>
	   {
	       var selected = Aws.Ec2.GetSubnets.Invoke(new()
	       {
	           Filters = new[]
	           {
	               new Aws.Ec2.Inputs.GetSubnetsFilterInputArgs
	               {
	                   Name = "tag:Name",
	                   Values = new[]
	                   {
	                       "",
	                   },
	               },
	           },
	       });

	   });
	   ```
	   ```go
	   package main

	   import (
	   	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	   	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	   )

	   func main() {
	   	pulumi.Run(func(ctx -pulumi.Context) error {
	   		_, err := ec2.GetSubnets(ctx, &ec2.GetSubnetsArgs{
	   			Filters: []ec2.GetSubnetsFilter{
	   				{
	   					Name: "tag:Name",
	   					Values: []string{
	   						"",
	   					},
	   				},
	   			},
	   		}, nil)
	   		if err != nil {
	   			return err
	   		}
	   		return nil
	   	})
	   }
	   ```
	   ```java
	   package generated_program;

	   import com.pulumi.Context;
	   import com.pulumi.Pulumi;
	   import com.pulumi.core.Output;
	   import com.pulumi.aws.ec2.Ec2Functions;
	   import com.pulumi.aws.ec2.inputs.GetSubnetsArgs;
	   import java.util.List;
	   import java.util.ArrayList;
	   import java.util.Map;
	   import java.io.File;
	   import java.nio.file.Files;
	   import java.nio.file.Paths;

	   public class App {
	       public static void main(String[] args) {
	           Pulumi.run(App::stack);
	       }

	       public static void stack(Context ctx) {
	           final var selected = Ec2Functions.getSubnets(GetSubnetsArgs.builder()
	               .filters(GetSubnetsFilterArgs.builder()
	                   .name("tag:Name")
	                   .values("")
	                   .build())
	               .build());

	       }
	   }
	   ```
	   ```yaml
	   variables:
	     selected:
	       fn::invoke:
	         Function: aws:ec2:getSubnets
	         Arguments:
	           filters:
	             - name: tag:Name
	               values:
	                 -
	   ```
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Set of values that are accepted for the given field.
	   Subnet IDs will be selected if any one of the given values match.
	*/
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
