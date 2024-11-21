package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/open-policy-agent/opa/rego"

	"rigSecurityMaytar/githubFetcher/proto"
)

const (
	policiesDirPath      = "../../policies/githubRepository"
	policyFileNameSuffix = ".rego"
)

type policyMetadata struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func evaluatePolicies(ctx context.Context, input map[string]interface{}) ([]*proto.Policy, error) {
	policiesResult := make([]*proto.Policy, 0)

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("[%s] error getting current directory: %w", ctx, err)
	}

	err = filepath.Walk(path.Join(currentDir, policiesDirPath), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, policyFileNameSuffix) {
			status, err := evaluatePolicy(ctx, path, input)
			if err != nil {
				return fmt.Errorf("[%s] error evaluating policy %s: %s", ctx, path, err)
			}

			metadata, err := getPolicyMetadata(ctx, path)
			if err != nil {
				return fmt.Errorf("[%s] error getting metadata for policy %s: %s", ctx, path, err)
			}

			policyResult := &proto.Policy{
				Id:     metadata.Id,
				Title:  metadata.Title,
				Result: status,
			}

			policiesResult = append(policiesResult, policyResult)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[%s] error walking through policies directories: %w", ctx, err)
	}

	return policiesResult, nil
}

func getPolicyMetadata(ctx context.Context, policyPath string) (*policyMetadata, error) {
	metadataFilePath := filepath.Join(filepath.Dir(policyPath), "metadata.json")

	metadataFile, err := os.Open(metadataFilePath)
	if err != nil {
		return nil, fmt.Errorf("[%s] error opening metadata file: %w", ctx, err)
	}
	defer metadataFile.Close()

	var metadata policyMetadata
	decoder := json.NewDecoder(metadataFile)
	if err = decoder.Decode(&metadata); err != nil {
		return nil, fmt.Errorf("[%s] error parsing metadata file: %w", ctx, err)
	}

	return &metadata, nil
}

// evaluatePolicy evaluates the policy against the input, it will return true if the policy returned pass == true
func evaluatePolicy(ctx context.Context, policyPath string, input map[string]interface{}) (bool, error) {
	ctx = context.WithValue(ctx, "policyPath", policyPath)

	query, err := rego.New(
		rego.Query("data.policies.pass"),
		rego.Load([]string{policyPath}, nil),
	).PrepareForEval(ctx)
	if err != nil {
		return false, fmt.Errorf("[%s] failed to prepare policy: %w", ctx, err)
	}

	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return false, fmt.Errorf("[%s] failed to evaluate policy: %w", ctx, err)
	}

	if len(results) == 0 {
		return true, nil
	}

	if len(results[0].Expressions) > 0 {
		pass, ok := results[0].Expressions[0].Value.(bool)
		if !ok {
			return false, fmt.Errorf("[%s] failed to evaluate policy result: %w", ctx, results[0].Expressions[0].Value)
		}

		return pass, nil
	}

	return true, nil
}
